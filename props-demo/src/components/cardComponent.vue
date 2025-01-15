<template>
<div class="card" v-for="data in cardData" :key="data.name">
    <img :src="data.image" alt="Card Image" class="card-image" />
    <div class="card-content">
        <h2 class="card-title">{{data.title}}</h2>
        <p class="card-subinfo">{{data.description}}</p>
        <button @click="getName(data.title,data.description)" class="card-button">Click Me</button>
    </div>
</div>
</template>

  
<script>
    export default {
        name:"CardComponent",
        props:{
            cardData:Object
        },
        methods : {
            getName(title,info){
                alert("Name : "+title+"\n"+info)
            }
        }
    }
</script>

  
<style scoped>
.card {
    width: 300px;
    border: 1px solid #ddd;
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    overflow: hidden;
    text-align: center;
    margin: 20px;
    background-color: #fff;
}

.card-image {
    width: 100%;
    height: 200px;
}

.card-content {
    padding: 16px;
}

.card-title {
    font-size: 18px;
    font-weight: bold;
    margin: 8px 0;
}

.card-subinfo {
    font-size: 14px;
    color: #666;
    margin: 8px 0;
}

.card-button {
    padding: 10px 20px;
    font-size: 14px;
    color: #fff;
    background-color: #007bff;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.3s ease;
}

.card-button:hover {
    background-color: #0056b3;
}
</style>
