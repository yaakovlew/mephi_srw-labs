import { createSlice } from '@reduxjs/toolkit';

const authSlice = createSlice({
    name: 'auth',
    initialState: {
        token: localStorage.getItem('token') || null,
        labToken: localStorage.getItem('lab-token') || null,
    },
    reducers: {
        setToken: (state, action) => {
            state.token = action.payload;
            localStorage.setItem('token', action.payload);
        },
        clearToken: (state) => {
            state.token = null;
            localStorage.removeItem('token');
        },
        setLabToken: (state, action) => {
            state.labToken = action.payload;
            localStorage.setItem('lab-token', action.payload);
        },
        clearLabToken: (state) => {
            state.labToken = null;
            localStorage.removeItem('lab-token');
        },
    },
});

export const {
    setToken,
    clearToken
} = authSlice.actions;
export default authSlice.reducer;
