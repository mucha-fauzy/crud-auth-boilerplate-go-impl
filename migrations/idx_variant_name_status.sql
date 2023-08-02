-- Composite Index
CREATE INDEX idx_variant_name_status ON variants (name, status);