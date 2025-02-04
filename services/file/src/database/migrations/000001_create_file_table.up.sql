CREATE TABLE IF NOT EXISTS FILES (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- UUID primary key, auto-generated by PostgreSQL
    fileUri TEXT NOT NULL,                         -- URI of the file stored in AWS S3
    thumbnailUri TEXT,                             -- URI of the thumbnail (nullable)
    created_at TIMESTAMP DEFAULT NOW(),            -- Automatically set creation timestamp
    updated_at TIMESTAMP DEFAULT NOW()             -- Automatically set update timestamp
);
