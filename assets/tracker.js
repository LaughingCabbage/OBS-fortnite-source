var app = new Vue({
    el: '#app',
    data: {
        timeout: 5000
    },
    created: function(){
        this.start();
    },
    methods:{
        start: function(){
            console.log(`starting`)
            setTimeout(function(){
                window.location.reload(1);
            }, this.timeout);
        }
    }
})