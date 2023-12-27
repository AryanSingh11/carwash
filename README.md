
# Service Station Management software in Golang        

   

<a href="https://ibb.co/r7ym0fX"><img src="https://i.ibb.co/8K80s4Q/carwash-Architectuere.png" alt="carwash-Architectuere" border="0"></a>      

or click here: 
https://ibb.co/r7ym0fX


There are 3 profiles and 2 reports and 1 email service available. This Readme would guide  you on how to use all these functionalities    

## how to run this locally?
step1 :  create a new folder, and in terminal run

` git clone https://github.com/AryanSingh11/carwash.git`  

step 2: change directory to carwash and in terminal run     
`cd carwash`         

`go run main.go `     

      
You'll get a beautiful terminal banner like this: 
<a href="https://ibb.co/YQ056jj"><img src="https://i.ibb.co/DWrBNKK/image1.png" alt="image1" border="0"></a>
      
      
## Profiles available are
### 1. Customer  : after step 2, type: customer to log into/access this profile
supports:   
 1. Register new Customer, type register to access this. You'd be promoted to enter all the data one by one and all the new cusotmer object will be generated.
 2. Create new Customer, type newcustomer to access this.      

3. Delete Customer, type delete to access this. You'd be prompted to enter customerID and customer obj corresponding to that ID will be deleted from customer.csv file      

4. Update Customer, type update to access this.You'd be prompted to enter customerID and customer obj corresponding to that ID will be updated in customer.csv file   

` note: all the data related to Customers is stored in customersNew.csv `

### 2. Dealer  : after step 2, type: dealer to log into/access this profile
these are the methods available for you as a Dealer
1. Register new Dealer, type registerdealer to access this. You'd be promoted to enter all the data one by one and all the new dealer object will be generated.       

2. Create new Dealer, type newdealer to access this
3. Delete Dealer, type deletedealer to access this. You'd be prompted to enter dealerID and dealer obj corresponding to that ID will be deleted from customer.csv file         

4. Update Dealer, type updatedealer to access this                

` note: all the data related to Dealer is stored in dealer.csv `       

### 3. Technician  : after step 2, type: technician to log into/access this profile
these are the methods available for you as a Technician.  

Works similar to Dealer    

1. Register new Technician, type registerTechnician to access this
2. Create new Technician, type newTechnician to access this
3. Delete Technician, type deleteTechnician to access this
3. Update Technician, type updateTechnician to access this  
` note: all the data related to technician is stored in technician.csv `            


## Reports available are
### 1. Report by date: type reportbydate to access this
1.  **Function**: Given a date, generates a report showing customer information, car, type of service and last date of service in the terminal     

2. after entering, enter the date in **YYYY-MM-DD** format to get the services information on this date       

`note: the report would be generated in the terminal itself`     

### 2. Notification Report: type notificationreport to access this
1. **Function**: 2.	Given customer id, generate a report showing details of last coupon notification for each of the service.       
        
        
**you have 2 methods available:**
1. **GetAllNotificationOnID: type getallnotificationonid to access this**: enter and ID and we will send list of all coupon notifications sent to this ID
2. **CreateNewNotificationOnID : type createnewnotificationonid to access this**: lets you create new notification for the CustomerID entered   
    
     
`note that method 2 would automatically capture today's date`

`note: the report would be generated in the terminal itself`           

      

## EMAIL SERVICE     
**Function** : This will send a promotion (or reminder) to the vehicle owner for oil change 6 months since the last oil change

**How to access** : To Run the Email service, type : email in terminal
 




