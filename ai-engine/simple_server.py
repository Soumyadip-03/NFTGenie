"""
Simplified FastAPI Server for NFTGenie AI Engine (No DB/Redis required)
"""

from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel, Field
from typing import List, Optional
from datetime import datetime
from contextlib import asynccontextmanager
import json

# Import our recommendation engine
from advanced_recommender import (
    AdvancedRecommendationEngine,
    NFT, User, Interaction
)

# Sample data for testing
sample_nfts = [
    NFT(
        id="1",
        name="Digital Art #001",
        description="Beautiful digital artwork",
        image_url="https://via.placeholder.com/300",
        creator_id="creator1",
        owner_id="owner1",
        tags=['art', 'digital'],
        attributes={'rarity': 'rare'},
        price=0.1,
        views=150,
        likes=25,
        created_at=datetime.now(),
        chain="polygonAmoy"
    ),
    NFT(
        id="2", 
        name="Gaming NFT #001",
        description="Rare gaming collectible",
        image_url="https://via.placeholder.com/300",
        creator_id="creator2",
        owner_id="owner2", 
        tags=['gaming', 'collectible'],
        attributes={'power': 100},
        price=0.5,
        views=300,
        likes=50,
        created_at=datetime.now(),
        chain="polygonAmoy"
    )
]

# Initialize recommendation engine
recommendation_engine = AdvancedRecommendationEngine()

@asynccontextmanager
async def lifespan(app: FastAPI):
    """Initialize the application"""
    print("Starting NFTGenie AI Engine...")
    
    # Create sample data
    sample_users = [
        User(
            id="user1",
            wallet_address="0x123",
            interaction_history=[],
            preferences={'price_range': {'min': 0, 'max': 1}},
            created_at=datetime.now()
        )
    ]
    
    sample_interactions = [
        Interaction(
            user_id="user1",
            nft_id="1",
            type="view",
            timestamp=datetime.now(),
            value=1.0
        )
    ]
    
    # Train model with sample data
    recommendation_engine.train(sample_users, sample_nfts, sample_interactions)
    print("AI Engine ready")
    yield
    print("Shutting down AI Engine")

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
    allow_origins=["http://localhost:3000", "http://localhost:8080"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Pydantic models
class RecommendationRequest(BaseModel):
    user_id: str
    limit: int = Field(default=10, ge=1, le=100)
    strategy: str = Field(default="hybrid")

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
    total_nfts: int

# API Endpoints
@app.get("/", response_model=HealthResponse)
async def health_check():
    """Health check endpoint"""
    return HealthResponse(
        status="healthy",
        model_version=recommendation_engine.model_version,
        total_nfts=len(sample_nfts)
    )

@app.post("/recommend", response_model=List[RecommendationResponse])
async def get_recommendations(request: RecommendationRequest):
    """Get personalized NFT recommendations"""
    
    # Get recommendations
    recommendations = recommendation_engine.recommend(
        user_id=request.user_id,
        nfts=sample_nfts,
        k=request.limit,
        strategy=request.strategy
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
    
    return response

@app.get("/trending")
async def get_trending_nfts(limit: int = 10):
    """Get trending NFTs"""
    
    recommendations = recommendation_engine.recommend(
        user_id="anonymous",
        nfts=sample_nfts,
        k=limit,
        strategy="trending"
    )
    
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

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(
        "simple_server:app",
        host="0.0.0.0",
        port=5000,
        reload=True,
        log_level="info"
    )