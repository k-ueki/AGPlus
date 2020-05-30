<template>
        <v-btn
                bottom
                color="pink"
                dark
                fab
                fixed
                right
                @click="dialog = !dialog"
        >
            <v-icon>mdi-plus</v-icon>
        </v-btn>
</template>

<script>
    export default {
        name:'Buttom',
    }
</script>
