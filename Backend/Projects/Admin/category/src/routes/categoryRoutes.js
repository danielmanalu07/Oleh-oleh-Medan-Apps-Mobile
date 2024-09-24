const express = require('express');
const router = express.Router();
const categoryController = require('../controllers/categoryController');
const authMiddleware = require('../middleware/auth');

router.post('/category/create', authMiddleware, categoryController.createCategory);

module.exports = router;