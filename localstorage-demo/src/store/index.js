import { createStore } from 'vuex';

export const store = createStore({
  state: {
    myData: localStorage.getItem('myData') || '', // Initialize from localStorage
  },
  mutations: {
    setMyData(state, data) {
      state.myData = data;
      localStorage.setItem('myData', data); // Persist to localStorage
    }
  },
  actions: {
    updateMyData({ commit }, data) {
      commit('setMyData', data);
    }
  }
});
