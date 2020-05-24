<template>
    <div>
    <div class="row">
        <div class="col col-md-10 offset-1">
                <div class="card">
                    <div class="card-body">
                    <div class="row">
                            <div class="col col-md-2">
                                <div class="box-image">
                                    <div v-if="domain.logo != ''">
                                        <img v-bind:src="domain.logo">
                                    </div>
                                    <div v-else>
                                        <img src="../assets/internet-logo.png">
                                    </div>
                                    
                                    
                                </div>
                            </div>
                            <div class="col col-md-6 align-self-center">
                                <h4 v-if="domain.title != ''" class="text-center">{{domain.title}}</h4>
                                <h4 v-else class="text-center">{{domain.name}}</h4>
                            </div>
                            <div class="col col-md-1 align-self-center">
                                <h4><span class="badge badge-success">{{domain.ssl_grade}}</span></h4>
                            </div>

                            <div class="col col-md-2">
                                <b-button v-b-toggle="domain.name" variant="link">Ver Detalles</b-button>
                            </div>
                        </div>
                    </div>
                    <b-collapse v-bind:id="domain.name" class="mt-2">
                        <div class="row">
                            <div class="col col-2">
                                 <div class="form-check">
                                    <input type="checkbox" class="form-check-input" v-bind:value="domain.is_down" disabled>
                                    <label class="form-check-label">Is Down?</label>
                             </div>
                            </div>
                            <div class="col col-3">
                                <div class="form-check">
                                    <input type="checkbox" class="form-check-input" v-bind:value="domain.servers_changed" disabled>
                                    <label class="form-check-label">Servers changed?</label>
                                </div>
                            </div>
                        </div>
                       
                         
                    <b-card>
                        
                        <table class="table table-hover">
                            <thead>
                                <tr>
                                <th scope="col">Address</th>
                                <th scope="col">Owner</th>
                                <th scope="col">Country</th>
                                <th scope="col">SSL Grade</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr 
                                    v-for="server in domain.servers"
                                    v-bind:server="server"
                                    v-bind:key="server.address">
                                    <td> {{server.address}}</td>
                                    <td class="text-capitalize"> {{server.owner}}</td>
                                    <td class="text-uppercase"> {{server.country}}</td>
                                    <td> <span class="badge badge-success">{{server.ssl_grade}}</span></td>
                                </tr>
                            </tbody>
                        </table>
                      
                    </b-card>
                    </b-collapse>
                
                </div>
        </div>
    </div>
    <br>
    </div>
</template>

<script>
export default {
    name: 'DoaminComponent',
    props: ['domain']
}
</script>
<style scoped>
.box-image {
    width: 50px;
    
}
img {
     max-width: 100%;
    max-height: 100%;
    display: block; /* remove extra space below image */
}
.div-centered {
  text-align: left; 
  margin: auto; 
  
}
</style>