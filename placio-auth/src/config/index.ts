import { config } from 'dotenv';
config({ path: `.env.${process.env.NODE_ENV || 'development'}.local` });

export const CREDENTIALS = process.env.CREDENTIALS === 'true';
export const { NODE_ENV, PORT, SECRET_KEY, LOG_FORMAT, LOG_DIR, ORIGIN, CRYPTO_SECRET } = process.env;
export const { DB_HOST, DB_PORT, DB_DATABASE, DB_USERNAME, DB_PASSWORD } = process.env;
export const { TWITTER_CONSUMER_KEY, TWITTER_CONSUMER_SECRET, FACEBOOK_APP_ID, FACEBOOK_APP_SECRET } = process.env;