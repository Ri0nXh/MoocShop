-- 商品表
CREATE TABLE `goods` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '商品名称',
    `pic_url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '商品图片',
    `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '商品价格',
    `stock` int NOT NULL DEFAULT '0' COMMENT '库存',
    `sale` int NOT NULL DEFAULT '0' COMMENT '销量',
    `brand` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '品牌',
    `tag` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '标签',
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    `deleted_at` datetime DEFAULT NULL,
    `detail` text COMMENT '商品详情',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;