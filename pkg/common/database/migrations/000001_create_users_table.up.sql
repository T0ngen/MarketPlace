CREATE TABLE "goods" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "seller_id" bigint NOT NULL,
  "title" varchar NOT NULL,
  "price" bigint NOT NULL,
  "description" varchar NOT NULL,
  "image" varchar NOT NULL,
  "category" varchar NOT NULL,
  "rating" varchar NOT NULL,
  "discount" bigint NOT NULL,
  "status" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON "goods" ("title");
INSERT INTO goods (seller_id,title, price, description, image, category, rating, discount, status) 
VALUES 
(1, 'Product1', 100, 'Description1', 'image1.jpg', 'Category1', '4.5', 0, 'Active'),
(2, 'Product2', 150, 'Description2', 'image2.jpg', 'Category2', '4.7', 10, 'Inactive'),
(3, 'Product3', 200, 'Description3', 'image3.jpg', 'Category1', '4.2', 5, 'Active'),
(4, 'Product4', 120, 'Description4', 'image4.jpg', 'Category3', '4.8', 15, 'Active'),
(5, 'Product5', 180, 'Description5', 'image5.jpg', 'Category2', '4.3', 8, 'Inactive'),
(6, 'Product6', 250, 'Description6', 'image6.jpg', 'Category3', '4.6', 12, 'Active'),
(7, 'Product7', 300, 'Description7', 'image7.jpg', 'Category1', '4.9', 20, 'Active'),
(8, 'Product8', 90, 'Description8', 'image8.jpg', 'Category2', '4.4', 5, 'Inactive'),
(9, 'Product9', 160, 'Description9', 'image9.jpg', 'Category3', '4.7', 8, 'Active'),
(10, 'Product10', 140, 'Description10', 'image10.jpg', 'Category1', '4.5', 10, 'Active');


INSERT INTO goods (seller_id, title, price, description, image, category, rating, discount, status) 
VALUES 
(1, 'Product11', 120, 'Description11', 'image11.jpg', 'Category3', '4.3', 5, 'Active'),
(2, 'Product12', 180, 'Description12', 'image12.jpg', 'Category1', '4.6', 8, 'Inactive'),
(3, 'Product13', 250, 'Description13', 'image13.jpg', 'Category2', '4.8', 12, 'Active'),
(4, 'Product14', 300, 'Description14', 'image14.jpg', 'Category1', '4.5', 20, 'Active'),
(5, 'Product15', 90, 'Description15', 'image15.jpg', 'Category3', '4.7', 5, 'Inactive'),
(6, 'Product16', 160, 'Description16', 'image16.jpg', 'Category2', '4.2', 8, 'Active'),
(7, 'Product17', 140, 'Description17', 'image17.jpg', 'Category3', '4.9', 10, 'Active'),
(8, 'Product18', 120, 'Description18', 'image18.jpg', 'Category1', '4.4', 5, 'Inactive'),
(9, 'Product19', 200, 'Description19', 'image19.jpg', 'Category2', '4.7', 10, 'Active'),
(10, 'Product20', 150, 'Description20', 'image20.jpg', 'Category1', '4.6', 12, 'Active');