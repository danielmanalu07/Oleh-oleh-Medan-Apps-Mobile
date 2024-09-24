const categoryService = require('../services/categoryService');
const getAdminData = require('../utils/auth_admin');

async function createCategory(req, res) {
    try {
        const { nama, image } = req.body;
        const token = req.headers.authorization;

        const adminData = await getAdminData(token);
        const adminId = adminData.data.id;

        const newCategory = await categoryService.createCategory({
            nama,
            image,
            admin_id: Number(adminId)
        });

        res.status(201).json(newCategory);
    } catch (error) {
        res.status(500).json({
            error: error.message,
        });
    }
}

module.exports = {
    createCategory,
};