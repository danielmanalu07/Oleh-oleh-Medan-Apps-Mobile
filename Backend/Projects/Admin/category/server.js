require('dotenv').config();
const express = require('express');
const app = express();
const categoryRoutes = require('./src/routes/categoryRoutes');

app.use(express.json());

app.use('/api', categoryRoutes);

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
    console.log(`server is running on port ${PORT}`);
});