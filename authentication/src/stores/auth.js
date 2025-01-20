import { defineStore } from 'pinia';

export const useUserAuth = defineStore('userAuthentication', {
  state: () => ({
    userData: [],
  }),

  actions: {
    userRegistration(username, password, email, hobbies, gender) {
      const user = {
        username,
        password,
        email,
        hobbies,
        gender,
      };
      this.userData.push(user);
    },
    
  },


  persist: {
    enabled: true,  
    strategies: [
      {
        storage: sessionStorage,  
        key: 'user-auth',
        paths : ['userData']
      },
    ],
  },
});
