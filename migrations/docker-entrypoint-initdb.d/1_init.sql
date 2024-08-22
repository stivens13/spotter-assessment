-- Create channels table
-- enforce the channel_id to be unique and fixed length of 24

CREATE TABLE IF NOT EXISTS channels(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    channel_id varchar(24) NOT NULL CHECK (char_length(channel_id) = 24),
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    PRIMARY KEY(id),
    CONSTRAINT uc_channels_channel_id UNIQUE(channel_id)
);
CREATE INDEX idx_channels_deleted_at ON channels USING btree (channel_id, deleted_at);

-- Create channels table
-- enforce the video_id to be unique and fixed length of 11
-- add foreign key constraint to channel_id
-- add index on channel_id and and upload_date for faster query
-- in reverse chronological order

CREATE TABLE videos (
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    video_id varchar(11) NOT NULL CHECK (char_length(video_id) = 11),
    channel_id varchar(24) NOT NULL CHECK (char_length(channel_id) = 24),
    video_title varchar(255),
    upload_date date,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    PRIMARY KEY(id),
    CONSTRAINT fk_videos_channel_id FOREIGN KEY(channel_id) REFERENCES channels(channel_id),
    CONSTRAINT uc_video UNIQUE(video_id) 
);

CREATE INDEX idx_videos_channel_id_deleted_at ON videos USING btree (channel_id, deleted_at);
CREATE INDEX idx_videos_upload_date_channel_id ON videos USING btree (channel_id, upload_date DESC NULLS LAST);

