const { PrismaClient } = require('@prisma/client');
const prisma = new PrismaClient();

async function createCategory(data) {
    const category = await prisma.category.create({
        data,
    });
    return category;
}

module.exports = {
    createCategory,
};