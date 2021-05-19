<template>
<v-app>
  <v-data-table
    :headers="Headers"
    :items="doctors"
    :items-per-page="10"
    align="center">
  <template v-slot:[`item.actions`]="{ item }">
            <v-icon small class="mr-2" @click="editDoctor(item.id)">mdi-pencil</v-icon>
            <v-icon small @click="deleteDoctor(item.id)">mdi-delete</v-icon>
            <v-icon small link-to="/patients">mdi-calendar</v-icon>
          </template>
        
  </v-data-table>
  <v-row justify="center">
    <v-btn
      color="primary"
      dark
      @click.stop="adddialog = true"
    >
      add Doctor
    </v-btn>
    

    <v-dialog
      v-model="adddialog"
      id="adddialog"
      name="adddialog"
           
    >
      <v-card>
        <v-card-title class="headline">
          Add doctor
        </v-card-title>

        <v-card-text>
          fill the doctor_form to add a doctor
        </v-card-text>
      <form class="form" v-on:submit.prevent="createDoctor">
        <v-text-field

                  id="name"
                  label="name"
                  v-model="doctor_form.name"
                  required
                >
        </v-text-field>
        <v-text-field
                  id="years_of_experience"
                  label="years_of_experience"
                  v-model="doctor_form.years_of_experience"
                  required
                >
        </v-text-field>
         <v-text-field
                  id="type_of_doctor"                  
                  label="type_of_doctor"
                  v-model="doctor_form.type_of_doctor"
                  required
                >
        </v-text-field>
         <v-text-field
                  id="consultation_fee"
                  label="consultation_fee"
                  v-model="doctor_form.consultation_fee"
                  required
                >
        </v-text-field>
         <v-text-field
                  id="mail_id"
                  label="mail_id"
                  v-model="doctor_form.mail_id"
                  required

                >
        </v-text-field>
         <v-text-field
                  id="phone"
                  label="phone"
                  v-model="doctor_form.phone"
                  required
                >
        </v-text-field>
        
        <v-card-actions>
          <v-spacer></v-spacer>

          

          <v-btn
            color="green darken-1"
            text
            
            @click="add_edit_Doctor()"
          >
            submit
          </v-btn>
          <v-btn
          @click="adddialog = false;
          id ='add'
            ">
            close
          </v-btn>

          
        </v-card-actions>
        
      </form>
      </v-card>
    </v-dialog>
   
  </v-row>
</v-app>
</template>
<script>
import doctorServices from "@/services/doctorService" 
import axios from 'axios'
axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*';
const Host="127.0.0.1"
const Port="5000"



  export default {
    data :() => ({

        doctor_form :{
          name:'',
          years_of_experience:'',
          type_of_doctor:'', 
          consultation_fee:'',
          mail_id:'',
          phone:'',
          hospital_id:'55555',


        }
      ,
     
      
      adddialog: false,
 

        Headers:[
            {text:'doctor id', value:"id"},
            {text:'doctor name', value:"name"},
            {text:'type of doctor', value:"type_of_doctor"},
            {text:'years of experience', value:"years_of_experience"},
            {text:'consultation fee', value:"consultation_fee"},
            {text:'email_id', value:"mail_id"},
            {text:'phone numbers', value:"phone"},
            {text:'actions',value:"actions"},
            

            
        ],
        doctors:[] ,
        
        
        

    }),
    created(){
        console.log('inside created')
        this.initialize()
    },
    methods:{
        initialize()
        {this.loadDoctors()},
        createDoctor(){
          this.doctor_form={
          name:'',
          years_of_experience:'',
          type_of_doctor:'', 
          consultation_fee:'',
          mail_id:'',
          phone:'',
          hospital_id:'55555',


        }
          this.adddialog= true
          
        },
        
        
        



        loadDoctors(){

            this.doctors=[]
            doctorServices.getAlldoctors()
            doctorServices.getAlldoctors().then(item=>{
                    item.doctor.forEach(
                        doctor=>{this.doctors.push(doctor)}
                    )
                    
                

            });
            

        },
        add_edit_Doctor(){
          if (this.doctor_form.id==null){

          
          axios.post(`'http://${Host}:${Port}/doctors'`, this.doctor_form)
          .then((res)=>{
            console.log(res);
            if (res.status==200){
              this.adddialog = false
              this.loadDoctors()
            }
          })
          .catch((err)=>{
            console.log(err);
          })
          .finally(()=>{})}
          else{
            axios.put(`http://${Host}:${Port}/doctors/${this.doctor_form.id}`,this.doctor_form)
          }

        },

        editDoctor(id){
          axios.get(`http://${Host}:${Port}/doctors/${id}`)
          .then((res)=>{
             console.log(res.data.doctor);
             this.doctor_form=res.data.doctor
             this.adddialog = true
          }
)
          .catch(err => console.log(err))
         
          
        },
        deleteDoctor(id){
          axios.delete(`http://${Host}:${Port}/doctors/${id}`)
          this.loadDoctors()
        }
      
        
    }}
  
</script>

  