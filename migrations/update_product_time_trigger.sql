CREATE TRIGGER update_product_time_trigger
AFTER UPDATE ON brands
FOR EACH ROW
  UPDATE products 
  SET updated_at = NOW()  
  WHERE products.brand_id = NEW.id;