const TOKEN_KEY = 'token';

function _setToken(token) {
  localStorage.setItem(TOKEN_KEY, token);
}

function _getToken() {
  return localStorage.getItem(TOKEN_KEY);
}

function _removeToken() {
  localStorage.removeItem(TOKEN_KEY);
}

export {
  _setToken,
  _getToken,
  _removeToken
};
