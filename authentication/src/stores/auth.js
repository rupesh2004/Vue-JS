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

      console.log('User Registered:', this.userData); // Debug: Check the state
    },
  },

  persist: {
    enabled: true,
    strategies: [
      {
        key: 'user-auth',
        storage: sessionStorage, // Use sessionStorage
        paths: ['userData'], // Persist only `userData`
      },
    ],
  },
});
