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

-- 订单表
CREATE TABLE `order` (
     `id` int NOT NULL AUTO_INCREMENT,
     `number` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单编号',
     `user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
     `pay_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '支付方式 1微信 2支付宝 3云闪付',
     `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
     `pay_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '支付时间',
     `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价',
     `consignee_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收货人姓名',
     `consignee_phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收货人手机号',
     `consignee_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收货人详细地址',
     `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '订单金额 单位元',
     `coupon_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '优惠券金额 单位元',
     `actual_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '实际支付金额 单位元',
     `created_at` datetime DEFAULT NULL,
     `updated_at` datetime DEFAULT NULL,
     PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1201 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='订单表';

-- 订单商品关系表
CREATE TABLE `order_goods_ref` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT '商品维度的订单表',
    `order_id` int NOT NULL DEFAULT '0' COMMENT '关联的主订单表',
    `goods_id` int NOT NULL DEFAULT '0' COMMENT '商品id',
    `goods_options_id` int DEFAULT '0' COMMENT '商品规格id sku id',
    `count` int NOT NULL COMMENT '商品数量',
    `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
    `price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '订单金额 单位元',
    `coupon_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '优惠券金额 单位元',
    `actual_price` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '实际支付金额 单位元',
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1748 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='订单商品关系表';

