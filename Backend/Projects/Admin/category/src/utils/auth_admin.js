const axios = require('axios');

async function getAdminData(token) {
    try {
        const response = await axios.post('http://localhost:8080/api/admin/login', {}, {
            headers: { Authorization: token }
        });
        return response.data;
    } catch (error) {
        throw new Error('Error fetching admin data');
    }
}

module.exports = getAdminData;