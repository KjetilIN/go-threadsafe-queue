# Q1: The Roller Coaster problem

## Problem definition: 

A number of passengers (1:n) wish to repeatedly take rides with a Roller Coaster. There should be one process for each passengers and one for the Roller Coaster. As for the problem of the sleeping barber, the processes must participate in several synchronization stages. The passengers must wait for the car to be present, and the car must wait for **C** passengers to enter. After the ride, the car must awake all riding passengers. The solution does not have to prevent sneaking. Make a solution in the setting of monitors with signal and continue discipline for signaling. Use the await language to solve this exercise. In addition, try to formulate a reasonable invariant based on your solution, but it is not necessary to prove the invariant using programming logic. 

## Solution:

- Passengers must wait until cart is available 
- Cart must be full before leaving 
- Cart must wake all passengers  
- Signal and continue 

```text
monitor RollerCoaster{
    // Max 
    int C; 

    int waitingPassengers = 0;         
    int passengersOnBoard = 0;   
    int rollerCoaster = 0;     
    int passenger_count = 0;                               

    condition carAvailable  
    condition carLeft      

    cond cart_occupied;
    cond cart_full;

    condition allPassengersBoarded      
    condition rideFinished   


    procedure ride(){
        // Wait until cart is here 
        while (rollerCoaster = 0){ wait(carAvailable) };

        // Passenger enter the 
        passengerOnBoard = passengerOnBoard + 1;
        signal(cart_occupied);

        // Wait until cart is full
        while(passengerOnBoard < C){ wait(cart_full); }

        // CART GOES FOR RIDE
        
        // Wait until cart is back
        while (rollerCoaster = 0){ wait(carAvailable) };

        // Leave passenger 
        passengerOnBoard = passengerOnBoard - 1; 
    }           

    procedure cartArrives(){
        while (passengersOnBoard = capacity){ wait(allPassengersBoarded) };
        rollerCoaster = rollerCoaster - 1;
        signal_all(carLeft); 
    }

    procedure passengerLeaves(){
        ...
    }


    // 1..N passengers call this 
    procedure passengerEnter(){
        # Increment the queue
        waitingPassengers += 1;

        wait(cartAvailable);


        passengersOnBoard += 1;
        waitingPassengers -= 1;

        if (passengersOnBoard = capacity){
            signal(allPassengersBoarded);
        }

        wait(rideFinished);
    }

}
```