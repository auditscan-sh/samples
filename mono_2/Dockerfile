# INTENTIONALLY WASTEFUL DOCKERFILE FOR COST/PERF DEMO
FROM python:3.11

# Wasteful: installs unused heavy dependencies
RUN apt-get update && apt-get install -y \
    vim \
    curl \
    wget \
    git \
    build-essential \
    libgl1-mesa-glx \
    ffmpeg \
    imagemagick \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY requirements.txt .
# No layer caching, copies everything, installs with no --no-cache-dir is waste but we add extra bloat
RUN pip install --no-cache-dir -r requirements.txt
RUN pip install torch torchvision tensorflow --no-cache-dir

COPY . .

# Wasteful: runs as root, no HEALTHCHECK, large image
CMD ["python", "app.py"]
