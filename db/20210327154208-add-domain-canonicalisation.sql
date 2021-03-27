CREATE OR REPLACE FUNCTION canon(s text) RETURNS text AS $$
	BEGIN
		RETURN UPPER(regexp_replace(s, '[.:]', ''));
	END;
$$ LANGUAGE plpgsql;
