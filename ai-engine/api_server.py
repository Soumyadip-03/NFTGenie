"""
FastAPI Server for NFTGenie AI Recommendation Engine
Provides REST API endpoints for NFT recommendations
"""

from fastapi import FastAPI, HTTPException, Depends, BackgroundTasks
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel, Field
from typing import List, Optional, Dict, Any
from datetime import datetime, timedelta
from contextlib import asynccontextmanager, contextmanager
import os
import json
import asyncio
from loguru import logger
import psycopg2
from psycopg2.extras import RealDictCursor
import redis
from dotenv import load_dotenv

# Import our recommendation engine
from advanced_recommender import (
    AdvancedRecommendationEngine,
    NFT, User, Interaction
)

# Load environment variables
load_dotenv()

@asynccontextmanager
async def lifespan(app: FastAPI):
    """Initialize and cleanup the application"""
    logger.info("Starting NFTGenie AI Engine...")
    
    # Try to load existing model
    model_path = "models/recommendation_model.json"
    if os.path.exists(model_path):
        try:
            recommendation_engine.load_model(model_path)
            logger.info("Loaded existing model")
        except Exception as e:
            logger.error(f"Failed to load model: {e}")
            await train_model()
    else:
        await train_model()
    
    logger.info("AI Engine ready")
    yield
    logger.info("Shutting down AI Engine")

# Initialize FastAPI app
app = FastAPI(
    title="NFTGenie AI Engine",
    description="AI-powered NFT recommendation service",
    version="1.0.0",
    lifespan=lifespan
)

# CORS configuration
app.add_middleware(
    CORSMiddleware,
    allow_origins=os.getenv("ALLOWED_ORIGINS", "http://localhost:3000").split(","),
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Initialize recommendation engine
recommendation_engine = AdvancedRecommendationEngine()

# Redis client for caching
redis_client = redis.Redis(
    host=os.getenv("REDIS_HOST", "localhost"),
    port=int(os.getenv("REDIS_PORT", 6379)),
    db=int(os.getenv("REDIS_DB", 0)),
    decode_responses=True
)

# Database configuration
DB_CONFIG = {
    "host": os.getenv("DB_HOST", "localhost"),
    "port": os.getenv("DB_PORT", "5432"),
    "database": os.getenv("DB_NAME", "nftgenie"),
    "user": os.getenv("DB_USER", "postgres"),
    "password": os.getenv("DB_PASSWORD", "")
}


# Pydantic models for API
class RecommendationRequest(BaseModel):
    user_id: str
    limit: int = Field(default=10, ge=1, le=100)
    strategy: str = Field(default="hybrid", pattern="^(hybrid|collaborative|content|trending)$")
    exclude_owned: bool = True
    diversify: bool = True


class InteractionData(BaseModel):
    user_id: str
    nft_id: str
    interaction_type: str = Field(pattern="^(view|like|purchase|mint)$")
    value: float = Field(default=1.0, ge=0, le=10)


class TrainRequest(BaseModel):
    force_retrain: bool = False


class RecommendationResponse(BaseModel):
    nft_id: str
    name: str
    image_url: str
    score: float
    reason: str
    price: float
    creator: str
    tags: List[str]


class HealthResponse(BaseModel):
    status: str
    model_version: str
    last_trained: Optional[str]
    cache_status: str


# Database connection
@contextmanager
def get_db():
    """Get database connection"""
    conn = psycopg2.connect(**DB_CONFIG)
    try:
        yield conn
    finally:
        conn.close()


# Helper functions
async def fetch_nfts_from_db() -> List[NFT]:
    """Fetch NFTs from database"""
    with get_db() as conn:
        with conn.cursor(cursor_factory=RealDictCursor) as cur:
            cur.execute("""
                SELECT n.*, u.wallet_address as creator_address
                FROM nfts n
                LEFT JOIN users u ON n.creator_id = u.id
                ORDER BY n.created_at DESC
                LIMIT 1000
            """)
            rows = cur.fetchall()
            
            nfts = []
            for row in rows:
                nft = NFT(
                    id=str(row['id']),
                    name=row['name'],
                    description=row.get('description', ''),
                    image_url=row['image_url'],
                    creator_id=str(row['creator_id']),
                    owner_id=str(row['owner_id']),
                    tags=row.get('tags', []),
                    attributes=json.loads(row['attributes']) if row.get('attributes') else {},
                    price=float(row.get('price', 0)) if row.get('price') else 0,
                    views=row['views'],
                    likes=row['likes'],
                    created_at=row['created_at'],
                    chain=row['chain'],
                    collection_id=str(row['collection_id']) if row.get('collection_id') else None
                )
                nfts.append(nft)
                
    return nfts


async def fetch_users_from_db() -> List[User]:
    """Fetch users from database"""
    with get_db() as conn:
        with conn.cursor(cursor_factory=RealDictCursor) as cur:
            cur.execute("""
                SELECT u.*, up.preferred_categories, up.preferred_price_range
                FROM users u
                LEFT JOIN user_preferences up ON u.id = up.user_id
                LIMIT 1000
            """)
            rows = cur.fetchall()
            
            users = []
            for row in rows:
                preferences = {}
                if row.get('preferred_price_range'):
                    preferences['price_range'] = json.loads(row['preferred_price_range'])
                if row.get('preferred_categories'):
                    preferences['categories'] = row['preferred_categories']
                    
                user = User(
                    id=str(row['id']),
                    wallet_address=row['wallet_address'],
                    interaction_history=[],
                    preferences=preferences,
                    created_at=row['created_at']
                )
                users.append(user)
                
    return users


async def fetch_interactions_from_db() -> List[Interaction]:
    """Fetch user interactions from database"""
    with get_db() as conn:
        with conn.cursor(cursor_factory=RealDictCursor) as cur:
            cur.execute("""
                SELECT * FROM user_interactions
                WHERE created_at > NOW() - INTERVAL '30 days'
                ORDER BY created_at DESC
                LIMIT 10000
            """)
            rows = cur.fetchall()
            
            interactions = []
            for row in rows:
                interaction = Interaction(
                    user_id=str(row['user_id']),
                    nft_id=str(row['nft_id']),
                    type=row['interaction_type'],
                    timestamp=row['created_at'],
                    value=row['interaction_value']
                )
                interactions.append(interaction)
                
    return interactions


async def train_model():
    """Train the recommendation model"""
    logger.info("Starting model training...")
    
    # Fetch data from database
    users = await fetch_users_from_db()
    nfts = await fetch_nfts_from_db()
    interactions = await fetch_interactions_from_db()
    
    # Train the model
    recommendation_engine.train(users, nfts, interactions)
    
    # Save model
    model_path = "models/recommendation_model.json"
    os.makedirs("models", exist_ok=True)
    recommendation_engine.save_model(model_path)
    
    # Update cache
    redis_client.set("model_last_trained", datetime.now().isoformat())
    
    logger.info("Model training completed")


# API Endpoints


@app.get("/", response_model=HealthResponse)
async def health_check():
    """Health check endpoint"""
    try:
        redis_client.ping()
        cache_status = "healthy"
    except:
        cache_status = "unavailable"
    
    last_trained = redis_client.get("model_last_trained")
    
    return HealthResponse(
        status="healthy",
        model_version=recommendation_engine.model_version,
        last_trained=last_trained,
        cache_status=cache_status
    )


@app.post("/recommend", response_model=List[RecommendationResponse])
async def get_recommendations(request: RecommendationRequest):
    """Get personalized NFT recommendations"""
    
    # Check cache
    cache_key = f"recommendations:{request.user_id}:{request.strategy}:{request.limit}"
    cached = redis_client.get(cache_key)
    
    if cached and not request.diversify:
        return json.loads(cached)
    
    # Fetch available NFTs
    nfts = await fetch_nfts_from_db()
    
    # Filter out owned NFTs if requested
    if request.exclude_owned:
        with get_db() as conn:
            with conn.cursor() as cur:
                cur.execute(
                    "SELECT id FROM nfts WHERE owner_id = (SELECT id FROM users WHERE wallet_address = %s)",
                    (request.user_id,)
                )
                owned_ids = {str(row[0]) for row in cur.fetchall()}
                nfts = [nft for nft in nfts if nft.id not in owned_ids]
    
    # Get recommendations
    try:
        recommendations = recommendation_engine.recommend(
            user_id=request.user_id,
            nfts=nfts,
            k=request.limit,
            strategy=request.strategy
        )
    except Exception as e:
        logger.error(f"Recommendation error: {e}")
        # Fallback to trending
        recommendations = recommendation_engine.recommend(
            user_id=request.user_id,
            nfts=nfts,
            k=request.limit,
            strategy="trending"
        )
    
    # Format response
    response = []
    for nft, score, reason in recommendations:
        response.append(RecommendationResponse(
            nft_id=nft.id,
            name=nft.name,
            image_url=nft.image_url,
            score=score,
            reason=reason,
            price=nft.price,
            creator=nft.creator_id,
            tags=nft.tags
        ))
    
    # Cache results
    redis_client.setex(
        cache_key,
        300,  # 5 minutes
        json.dumps([r.dict() for r in response])
    )
    
    return response


@app.post("/interaction")
async def record_interaction(interaction: InteractionData, background_tasks: BackgroundTasks):
    """Record user interaction and update model"""
    
    # Save to database
    with get_db() as conn:
        with conn.cursor() as cur:
            cur.execute("""
                INSERT INTO user_interactions (user_id, nft_id, interaction_type, interaction_value)
                VALUES (
                    (SELECT id FROM users WHERE wallet_address = %s),
                    %s::uuid,
                    %s,
                    %s
                )
                ON CONFLICT (user_id, nft_id, interaction_type) 
                DO UPDATE SET 
                    interaction_value = user_interactions.interaction_value + EXCLUDED.interaction_value,
                    created_at = NOW()
            """, (interaction.user_id, interaction.nft_id, 
                  interaction.interaction_type, interaction.value))
            conn.commit()
    
    # Update model in background
    background_tasks.add_task(
        update_model_with_interaction,
        interaction
    )
    
    # Clear cache for this user
    pattern = f"recommendations:{interaction.user_id}:*"
    for key in redis_client.scan_iter(match=pattern):
        redis_client.delete(key)
    
    return {"status": "success", "message": "Interaction recorded"}


async def update_model_with_interaction(interaction_data: InteractionData):
    """Update model with new interaction (background task)"""
    interaction = Interaction(
        user_id=interaction_data.user_id,
        nft_id=interaction_data.nft_id,
        type=interaction_data.interaction_type,
        timestamp=datetime.now(),
        value=interaction_data.value
    )
    recommendation_engine.update_real_time(interaction)


@app.post("/train")
async def train_model_endpoint(request: TrainRequest, background_tasks: BackgroundTasks):
    """Trigger model retraining"""
    
    # Check if training is needed
    if not request.force_retrain:
        last_trained = redis_client.get("model_last_trained")
        if last_trained:
            last_trained_time = datetime.fromisoformat(last_trained)
            if datetime.now() - last_trained_time < timedelta(hours=6):
                return {
                    "status": "skipped",
                    "message": "Model was trained recently",
                    "last_trained": last_trained
                }
    
    # Train in background
    background_tasks.add_task(train_model)
    
    return {
        "status": "training",
        "message": "Model training initiated in background"
    }


@app.get("/explain/{user_id}/{nft_id}")
async def explain_recommendation(user_id: str, nft_id: str):
    """Explain why an NFT was recommended to a user"""
    
    explanation = recommendation_engine.explain_recommendation(user_id, nft_id)
    
    return {
        "user_id": user_id,
        "nft_id": nft_id,
        "explanation": explanation
    }


@app.get("/trending")
async def get_trending_nfts(limit: int = 10):
    """Get trending NFTs"""
    
    # Fetch NFTs
    nfts = await fetch_nfts_from_db()
    
    # Get trending recommendations
    recommendations = recommendation_engine._trending_recommend(nfts, limit)
    
    response = []
    for nft, score, reason in recommendations:
        response.append({
            "nft_id": nft.id,
            "name": nft.name,
            "image_url": nft.image_url,
            "trend_score": score,
            "views": nft.views,
            "likes": nft.likes,
            "price": nft.price
        })
    
    return response


@app.get("/stats")
async def get_stats():
    """Get recommendation engine statistics"""
    
    with get_db() as conn:
        with conn.cursor(cursor_factory=RealDictCursor) as cur:
            # Get statistics
            cur.execute("""
                SELECT 
                    COUNT(DISTINCT user_id) as total_users,
                    COUNT(DISTINCT nft_id) as total_nfts,
                    COUNT(*) as total_interactions,
                    AVG(interaction_value) as avg_interaction_value
                FROM user_interactions
                WHERE created_at > NOW() - INTERVAL '30 days'
            """)
            stats = cur.fetchone()
    
    return {
        "model_version": recommendation_engine.model_version,
        "total_users": stats['total_users'],
        "total_nfts": stats['total_nfts'],
        "total_interactions": stats['total_interactions'],
        "avg_interaction_value": float(stats['avg_interaction_value'] or 0),
        "embeddings": {
            "users": len(recommendation_engine.user_embeddings),
            "nfts": len(recommendation_engine.nft_embeddings)
        },
        "cache_entries": redis_client.dbsize()
    }


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(
        "api_server:app",
        host="0.0.0.0",
        port=5000,
        reload=True,
        log_level="info"
    )
