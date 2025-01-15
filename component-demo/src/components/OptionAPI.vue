<script>
  export default {
    name:"OptionAPI",
    data(){
      return {
        name : "Rupesh Bhosale",
        email : "bhosalerupesh67@gmail.com",
        getMobileNum:()=>{
          return 8080186885;
        }
      }
    },
    methods : {
      getAddress (){
        return "Pune"
      },
      getAllData(){
        return {
          name : this.name,
          email : this.email,
          mobileNum : this.getMobileNum(),
          address : this.getAddress()
        }
      }
    }
  }
</script>

<template>
  <div class="optionAPI">
    <h3>Name : {{ getAllData().name }}</h3>
    <h3>Email : {{ email }}</h3>
    <h3>Mobile No. : {{ getMobileNum() }} </h3>
    <h3>Address : {{ getAddress() }}</h3>
    <h3>Add Data : {{ getAllData() }}</h3>
  </div>
</template>
