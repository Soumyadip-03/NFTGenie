# Simple content-based recommendation placeholder
# We'll implement a naive recommender using tags/keywords

from typing import List, Dict

# Example NFT data structure
# nft = {"id": "1", "name": "Art #1", "tags": ["art", "abstract"], "views": 10}


def recommend_nfts(user_history: List[Dict], nfts: List[Dict], k: int = 5) -> List[Dict]:
    """Return top-k NFTs by matching tags from user_history to available nfts.
    user_history: list of past interactions {"nft_id", "tags"}
    nfts: catalog with tags
    """
    # Build user interest score by tag
    tag_score = {}
    for h in user_history:
        for t in h.get("tags", []):
            tag_score[t] = tag_score.get(t, 0) + 1

    def score(nft):
        s = 0
        for t in nft.get("tags", []):
            s += tag_score.get(t, 0)
        # bonus for popularity
        s += 0.1 * nft.get("views", 0)
        return s

    ranked = sorted(nfts, key=score, reverse=True)
    return ranked[:k]

