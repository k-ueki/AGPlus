<template>
    <div class="side_bar">
        <href>Home</href><br/>
        <href>授業一覧</href><br/>
        <href>レビュー一覧</href><br/>
        <href>教員一覧</href><br/>
    </div>
</template>

<script>
    export default {
        name: 'ClassList',
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    h3 {
        margin: 40px 0 0;
    }
    ul {
        list-style-type: none;
        padding: 0;
    }
    li {
        display: inline-block;
        margin: 0 10px;
    }
    a {
        color: #000000;
    }
    .side_bar{
        float:left;
        background-color:#42b983;
    }
</style>
