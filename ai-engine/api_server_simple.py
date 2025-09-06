"""
Simplified FastAPI Server for NFTGenie AI Recommendation Engine
Works without Redis initially - uses in-memory caching
"""

from fastapi import FastAPI, HTTPException, BackgroundTasks
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel, Field
from typing import List, Optional, Dict, Any
from datetime import datetime, timedelta
import os
import json
from loguru import logger
from dotenv import load_dotenv

# Import our recommendation engine
from advanced_recommender import (
    AdvancedRecommendationEngine,
    NFT, User, Interaction
)

# Load environment variables
load_dotenv()

# Initialize FastAPI app
app = FastAPI(
    title="NFTGenie AI Engine",
    description="AI-powered NFT recommendation service",
    version="1.0.0"
)

# CORS configuration
app.add_middleware(
    CORSMiddleware,
    allow_origins=os.getenv("ALLOWED_ORIGINS", "http://localhost:3000,http://localhost:8000").split(","),
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Initialize recommendation engine
recommendation_engine = AdvancedRecommendationEngine()

# Simple in-memory cache (instead of Redis)
cache = {}

# Pydantic models for API
class RecommendationRequest(BaseModel):
    user_id: str
    limit: int = Field(default=10, ge=1, le=100)
    strategy: str = Field(default="hybrid", pattern="^(hybrid|collaborative|content|trending)$")
    exclude_owned: bool = True


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
    cache_entries: int


# Helper functions
def get_sample_nfts() -> List[NFT]:
    """Get sample NFTs for testing"""
    return [
        NFT(
            id="nft1",
            name="Digital Art #1",
            description="Beautiful digital artwork",
            image_url="https://via.placeholder.com/300",
            creator_id="creator1",
            owner_id="owner1",
            tags=['art', 'digital', 'abstract'],
            attributes={'rarity': 'rare'},
            price=10.0,
            views=150,
            likes=30,
            created_at=datetime.now() - timedelta(days=1),
            chain="polygonAmoy"
        ),
        NFT(
            id="nft2",
            name="Gaming NFT #2",
            description="Rare gaming collectible",
            image_url="https://via.placeholder.com/300",
            creator_id="creator2",
            owner_id="owner2",
            tags=['gaming', 'collectible'],
            attributes={'power': 100},
            price=25.0,
            views=500,
            likes=120,
            created_at=datetime.now() - timedelta(days=2),
            chain="polygonAmoy"
        ),
        NFT(
            id="nft3",
            name="Music NFT #3",
            description="Exclusive music track",
            image_url="https://via.placeholder.com/300",
            creator_id="creator3",
            owner_id="owner3",
            tags=['music', 'exclusive'],
            attributes={'duration': '3:45'},
            price=15.0,
            views=200,
            likes=45,
            created_at=datetime.now() - timedelta(days=3),
            chain="polygonAmoy"
        ),
        NFT(
            id="nft4",
            name="Photography NFT #4",
            description="Stunning landscape photo",
            image_url="https://via.placeholder.com/300",
            creator_id="creator1",
            owner_id="owner4",
            tags=['photography', 'landscape', 'art'],
            attributes={'resolution': '4K'},
            price=8.0,
            views=100,
            likes=20,
            created_at=datetime.now() - timedelta(days=4),
            chain="polygonAmoy"
        ),
        NFT(
            id="nft5",
            name="3D Model NFT #5",
            description="Interactive 3D model",
            image_url="https://via.placeholder.com/300",
            creator_id="creator4",
            owner_id="owner5",
            tags=['3d', 'interactive', 'digital'],
            attributes={'polygons': 50000},
            price=30.0,
            views=300,
            likes=60,
            created_at=datetime.now() - timedelta(hours=12),
            chain="polygonAmoy"
        )
    ]


def get_sample_users() -> List[User]:
    """Get sample users for testing"""
    return [
        User(
            id="user1",
            wallet_address="0x123...abc",
            interaction_history=[],
            preferences={'price_range': {'min': 0, 'max': 50}},
            created_at=datetime.now() - timedelta(days=30)
        ),
        User(
            id="user2",
            wallet_address="0x456...def",
            interaction_history=[],
            preferences={'price_range': {'min': 10, 'max': 100}},
            created_at=datetime.now() - timedelta(days=20)
        )
    ]


def get_sample_interactions() -> List[Interaction]:
    """Get sample interactions for testing"""
    return [
        Interaction(
            user_id="user1",
            nft_id="nft1",
            type="view",
            timestamp=datetime.now() - timedelta(hours=2),
            value=1.0
        ),
        Interaction(
            user_id="user1",
            nft_id="nft2",
            type="like",
            timestamp=datetime.now() - timedelta(hours=1),
            value=2.0
        ),
        Interaction(
            user_id="user2",
            nft_id="nft3",
            type="view",
            timestamp=datetime.now() - timedelta(hours=3),
            value=1.0
        )
    ]


# API Endpoints
@app.on_event("startup")
async def startup_event():
    """Initialize the application"""
    logger.info("Starting NFTGenie AI Engine (Simplified Mode)...")
    
    # Train model with sample data
    users = get_sample_users()
    nfts = get_sample_nfts()
    interactions = get_sample_interactions()
    
    recommendation_engine.train(users, nfts, interactions)
    
    logger.info("AI Engine ready (running without database/Redis)")


@app.get("/", response_model=HealthResponse)
async def health_check():
    """Health check endpoint"""
    return HealthResponse(
        status="healthy",
        model_version=recommendation_engine.model_version,
        cache_entries=len(cache)
    )


@app.post("/recommend", response_model=List[RecommendationResponse])
async def get_recommendations(request: RecommendationRequest):
    """Get personalized NFT recommendations"""
    
    # Check in-memory cache
    cache_key = f"recommendations:{request.user_id}:{request.strategy}:{request.limit}"
    
    if cache_key in cache:
        cached_data = cache[cache_key]
        # Simple TTL check (5 minutes)
        if datetime.now() - cached_data['timestamp'] < timedelta(minutes=5):
            return cached_data['data']
    
    # Get sample NFTs
    nfts = get_sample_nfts()
    
    # Filter out owned NFTs if requested (simplified)
    if request.exclude_owned and request.user_id == "user1":
        nfts = [nft for nft in nfts if nft.owner_id != "owner1"]
    
    # Get recommendations
    try:
        recommendations = recommendation_engine.recommend(
            user_id=request.user_id if request.user_id in ["user1", "user2"] else "user1",
            nfts=nfts,
            k=min(request.limit, len(nfts)),
            strategy=request.strategy
        )
    except Exception as e:
        logger.error(f"Recommendation error: {e}")
        # Fallback to trending
        recommendations = recommendation_engine._trending_recommend(nfts, request.limit)
    
    # Format response
    response = []
    for nft, score, reason in recommendations:
        response.append(RecommendationResponse(
            nft_id=nft.id,
            name=nft.name,
            image_url=nft.image_url,
            score=round(score, 3),
            reason=reason,
            price=nft.price,
            creator=nft.creator_id,
            tags=nft.tags
        ))
    
    # Cache results
    cache[cache_key] = {
        'timestamp': datetime.now(),
        'data': response
    }
    
    return response


@app.get("/trending")
async def get_trending_nfts(limit: int = 5):
    """Get trending NFTs"""
    
    nfts = get_sample_nfts()
    recommendations = recommendation_engine._trending_recommend(nfts, min(limit, len(nfts)))
    
    response = []
    for nft, score, _ in recommendations:
        response.append({
            "nft_id": nft.id,
            "name": nft.name,
            "image_url": nft.image_url,
            "trend_score": round(score, 3),
            "views": nft.views,
            "likes": nft.likes,
            "price": nft.price
        })
    
    return response


@app.get("/stats")
async def get_stats():
    """Get recommendation engine statistics"""
    return {
        "status": "running",
        "mode": "simplified (no database)",
        "model_version": recommendation_engine.model_version,
        "total_users": len(get_sample_users()),
        "total_nfts": len(get_sample_nfts()),
        "total_interactions": len(get_sample_interactions()),
        "embeddings": {
            "users": len(recommendation_engine.user_embeddings),
            "nfts": len(recommendation_engine.nft_embeddings)
        },
        "cache_entries": len(cache)
    }


@app.post("/clear-cache")
async def clear_cache():
    """Clear the in-memory cache"""
    cache.clear()
    return {"status": "success", "message": "Cache cleared"}


if __name__ == "__main__":
    import uvicorn
    logger.info("Starting server on http://localhost:5000")
    logger.info("API documentation available at http://localhost:5000/docs")
    uvicorn.run(
        app,
        host="0.0.0.0",
        port=5000,
        reload=False,
        log_level="info"
    )
