SELECT prod.Name as Product_name, cat.Name as Category_name FROM Products prod
LEFT JOIN Category_Produkt cat_prod on cat_prod.IDProd = prod.ID
LEFT JOIN Сategory cat on cat.ID = cat_prod.IDCat