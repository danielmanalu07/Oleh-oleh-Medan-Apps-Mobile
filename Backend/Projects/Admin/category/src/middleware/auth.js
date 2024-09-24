const authMiddleware = (req, res, next) => {
    const token = req.headers.authorization;
    if (!token) {
        return res.status(403).json({ error: 'Token is required' });
    }
    next();
};

module.exports = authMiddleware;
