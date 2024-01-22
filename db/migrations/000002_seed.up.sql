-- ActiveCampaign
--
-- ID: 1
INSERT INTO apps (name, slug)
    VALUES ('ActiveCampaign', 'activecampaign');

-- ID: 1
INSERT INTO app_fields (app_id, field)
    VALUES (1, 'api_key');

-- ID: 2
INSERT INTO app_fields (app_id, field)
    VALUES (1, 'api_url');

-- ID: 1
INSERT INTO integrations (name, app_id)
    VALUES ('ActiveCampaign', 1);

-- ActiveCampaign - API Key
INSERT INTO integration_app_field_values (integration_id, app_field_id, value)
    VALUES (1, 1, '');

-- ActiveCampaign - API URL
INSERT INTO integration_app_field_values (integration_id, app_field_id, value)
    VALUES (1, 2, '');

-- ConvertKit
--
-- ID: 2
INSERT INTO apps (name, slug)
    VALUES ('ConvertKit', 'convertkit');

-- ID: 3
INSERT INTO app_fields (app_id, field)
    VALUES (2, 'api_key');

-- ID: 4
INSERT INTO app_fields (app_id, field)
    VALUES (2, 'api_secret');

-- ID: 2
INSERT INTO integrations (name, app_id)
    VALUES ('ConvertKit', 2);

-- ConvertKit - API Key
INSERT INTO integration_app_field_values (integration_id, app_field_id, value)
    VALUES (2, 3, '');

-- ConvertKit - API Secret
INSERT INTO integration_app_field_values (integration_id, app_field_id, value)
    VALUES (2, 4, '');

