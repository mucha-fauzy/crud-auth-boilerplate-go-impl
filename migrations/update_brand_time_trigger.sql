CREATE TRIGGER update_brand_time_trigger
AFTER UPDATE ON variants
FOR EACH ROW
  UPDATE brands 
  SET updated_at = NOW()  
  WHERE brands.variant_id = NEW.id;