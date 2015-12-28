[![sherpa logo](http://nano-assets.gopagoda.io/readme-headers/sherpa.png)](http://nanobox.io/open-source#sherpa)
[![Build Status](https://travis-ci.org/nanopack/sherpa.svg)](https://travis-ci.org/nanopack/sherpa)

# sherpa

An api-driven approach to building machine images with Packer.

NOTICE: As of 12/28/15 this project has been deprecated and will no longer be
actively worked on. As it stands right now, sherpa can connect to a postgres
database and create/destroy templates and builds via an API.

There is an accompanying CLI that is mostly complete that can interface with the
API to accomplish the above. Otherwise some web interface or curl can be used.

The last piece that remains to be implemented is the actual integration with
Packer.

## Smaple DB Seed
```
CREATE TABLE templates (
created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
, download varchar NOT NULL
, template_id serial PRIMARY KEY
, transform_script varchar NOT NULL
);

CREATE TABLE builds (
created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()˜
, build_id serial PRIMARY KEY
, meta_data varchar NOT NULL
, state varchar NOT NULL DEFAULT 'incomplete'
, status varchar NOT NULL DEFAULT 'created'
, template_id int REFERENCES templates (template_id) ON UPDATE CASCADE
, transform_payload text NOT NULL
, updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE build_logs (
build_log_id serial PRIMARY KEY
, created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
, build_id int REFERENCES builds (build_id) ON UPDATE CASCADE ON DELETE CASCADE
, message text NOT NULL
, updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
```

## Status

Incomplete/Depricated - After further considerations, this project was deemed
unnecessary for our immediate needs. We may or may not come back and finish it
in the future, but the most likely scenario is not.

## Todo

- Documentation
- Tests

### Notes

[![sherpa logo](http://nano-assets.gopagoda.io/open-src/nanobox-open-src.png)](http://nanobox.io/open-source)
