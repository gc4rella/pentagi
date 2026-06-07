-- +goose Up
-- +goose StatementBegin
-- Add OpenRouter to the provider_type enum.
CREATE TYPE PROVIDER_TYPE_NEW AS ENUM (
  'openai',
  'anthropic',
  'gemini',
  'bedrock',
  'ollama',
  'custom',
  'deepseek',
  'glm',
  'kimi',
  'qwen',
  'openrouter'
);

ALTER TABLE providers
    ALTER COLUMN type TYPE PROVIDER_TYPE_NEW USING type::text::PROVIDER_TYPE_NEW;

ALTER TABLE flows
    ALTER COLUMN model_provider_type TYPE PROVIDER_TYPE_NEW USING model_provider_type::text::PROVIDER_TYPE_NEW;

ALTER TABLE assistants
    ALTER COLUMN model_provider_type TYPE PROVIDER_TYPE_NEW USING model_provider_type::text::PROVIDER_TYPE_NEW;

DROP TYPE PROVIDER_TYPE;
ALTER TYPE PROVIDER_TYPE_NEW RENAME TO PROVIDER_TYPE;

ALTER TABLE providers
    ALTER COLUMN type SET NOT NULL;

ALTER TABLE flows
    ALTER COLUMN model_provider_type SET NOT NULL;

ALTER TABLE assistants
    ALTER COLUMN model_provider_type SET NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM providers WHERE type = 'openrouter';
DELETE FROM flows WHERE model_provider_type = 'openrouter';
DELETE FROM assistants WHERE model_provider_type = 'openrouter';

CREATE TYPE PROVIDER_TYPE_NEW AS ENUM (
  'openai',
  'anthropic',
  'gemini',
  'bedrock',
  'ollama',
  'custom',
  'deepseek',
  'glm',
  'kimi',
  'qwen'
);

ALTER TABLE providers
    ALTER COLUMN type TYPE PROVIDER_TYPE_NEW USING type::text::PROVIDER_TYPE_NEW;

ALTER TABLE flows
    ALTER COLUMN model_provider_type TYPE PROVIDER_TYPE_NEW USING model_provider_type::text::PROVIDER_TYPE_NEW;

ALTER TABLE assistants
    ALTER COLUMN model_provider_type TYPE PROVIDER_TYPE_NEW USING model_provider_type::text::PROVIDER_TYPE_NEW;

DROP TYPE PROVIDER_TYPE;
ALTER TYPE PROVIDER_TYPE_NEW RENAME TO PROVIDER_TYPE;

ALTER TABLE providers
    ALTER COLUMN type SET NOT NULL;

ALTER TABLE flows
    ALTER COLUMN model_provider_type SET NOT NULL;

ALTER TABLE assistants
    ALTER COLUMN model_provider_type SET NOT NULL;
-- +goose StatementEnd
