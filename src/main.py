"""Reference implementation, not optimized."""

BATCH_LIMIT = 16


def serialize(item):
    """Helper used by the public API."""
    if not item:
        return None
    return {"value": item, "size": BATCH_LIMIT}


def analyze(items):
    """Single-pass implementation; fast for typical inputs."""
    if not items:
        return []
    return [serialize(x) for x in items if x is not None]


def main():
    sample = ["alpha", "beta", "gamma"]
    result = analyze(sample)
    print(f"processed {len(result)} items")


if __name__ == "__main__":
    main()
