import {createSlice} from "@reduxjs/toolkit";
import {_getToken, _removeToken, _setToken, request} from "../../utils";

const userStore = createSlice({
  name: "user",
  initialState: {
    token: _getToken() || '',
    userInfo: {}
  },
  reducers: {
    setToken(state, action) {
      state.token = action.payload;
      // 将 token 存储到 localStorage
      _setToken(action.payload);
    },
    setUserInfo(state, action) {
      state.userInfo = action.payload;
    },
    clearUserInfo(state) {
      state.token = '';
      state.userInfo = {};
      _removeToken();
    }
  }
});

const {setToken, setUserInfo, clearUserInfo} = userStore.actions;
const userReducer = userStore.reducer;

// 用户登录
const fetchLogin = (loginForm) => {
  return async (dispatch) => {
    const res = await request.post('/api/user/login', loginForm);
    // 确保即使 res.data 是 null 或 undefined 解构也不会失败
    const { token } = res.data ?? {};
    if (token) {
      dispatch(setToken(token));
    }
  };
}

// 获取用户信息
const fetchUserInfo = () => {
  return async (dispatch) => {
    const res = await request.get('/api/user/info');
    dispatch(setUserInfo(res.data.user));
  };
}

export {fetchLogin, fetchUserInfo, clearUserInfo};
export default userReducer;
