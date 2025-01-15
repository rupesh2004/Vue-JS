<script>
export default {
    name: "EventHandling",
    data(){
        return{
            count : 0
        }
    },
    methods: {
        handleSingEvent: () => {
            alert("Single Event");
        },
        handleDoubleClick : ()=>{
            alert("Double Click")
        },
        chanageToRed(){
            document.getElementById("text").style.color = "red";
        },
        changeToBlack(){
            document.getElementById("text").style.color = "black";
        },
        incrementCnt(){
            this.count++;
        }
    }
}
</script>

<template>
<div class="eventHandling">
    <h1 @mouseover="chanageToRed" id="text" @mouseout="changeToBlack">Event Handling</h1>
    <button @click="handleSingEvent">Single Click</button>
    <button @dblclick="handleDoubleClick">Double Click</button>
    <button @click="incrementCnt">Count : {{ count }}</button>
    <h2>Count is : {{ count }}</h2>
</div>
</template>
