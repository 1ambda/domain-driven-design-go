<template>
    <div id="app">
        <el-container>
            <el-header>
                <app-navbar></app-navbar>
            </el-header>
            <el-main>
                <router-view></router-view>
            </el-main>
        </el-container>
    </div>
</template>

<script lang="ts">
    import { Component, Vue } from "vue-property-decorator"
    import Navbar from "@/components/Navbar.vue" // @ is an alias to /src

    @Component({ components: { "app-navbar": Navbar } })
    export default class App extends Vue {

    }
</script>

<style lang="scss">
    @import './App.scss';
</style>
