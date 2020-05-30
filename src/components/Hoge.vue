<template>
        <v-content>
            <v-container
                    class="fill-height"
                    fluid
            >
                <v-row
                        align="center"
                        justify="center"
                >
                    <v-tooltip right>
                        <template v-slot:activator="{ on }">
                            <v-btn
                                    :href="source"
                                    icon
                                    large
                                    target="_blank"
                                    v-on="on"
                            >
                                <v-icon large>mdi-code-tags</v-icon>
                            </v-btn>
                        </template>
                        <span>Source</span>
                    </v-tooltip>
                </v-row>
            </v-container>
        </v-content>
</template>

<script>
    export default {
        name:'App',
        props: {
            source: String,
        },
        data: () => ({
            dialog: false,
            drawer: null,
            items: [
                {
                    icon: 'mdi-chevron-up',
                    'icon-alt': 'mdi-chevron-down',
                    text: '学部・学科',
                    model: false,
                    // children:[
                    //     {
                    //         text : '文学部',
                    //         model: false,
                    //         children:[
                    //             { text : '英米文学科'},
                    //             { text : '仏文学科'}
                    //         ],
                    //     },
                    // ],
                },
                { icon: 'mdi-message', text: 'Class Review' },
                { icon: 'mdi-history', text: 'My Favorite' },
                { icon: 'mdi-content-copy', text: '楽単ランキング' },
                // {
                //     icon: 'mdi-chevron-up',
                //     'icon-alt': 'mdi-chevron-down',
                //     text: 'Labels',
                //     model: true,
                //     children: [
                //         { icon: 'mdi-plus', text: 'Create label' },
                //     ],
                // },
                // {
                //     icon: 'mdi-chevron-up',
                //     'icon-alt': 'mdi-chevron-down',
                //     text: 'More',
                //     model: false,
                //     children: [
                //         { text: 'Import' },
                //         { text: 'Export' },
                //         { text: 'Print' },
                //         { text: 'Undo changes' },
                //         { text: 'Other contacts' },
                //     ],
                // },
                // { icon: 'mdi-settings', text: 'Settings' },
                // { icon: 'mdi-message', text: 'Send feedback' },
                // { icon: 'mdi-help-circle', text: 'Help' },
                // { icon: 'mdi-cellphone-link', text: 'App downloads' },
                // { icon: 'mdi-keyboard', text: 'Go to the old version' },
            ],
        }),
    }
</script>
