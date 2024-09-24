const prisma = require('../config/prisma');

async function createCategory(data) {
    return await prisma.category.create({ data });
}

module.exports = {
    createCategory,
};