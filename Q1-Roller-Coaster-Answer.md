# Q1: The Roller Coaster problem

## Problem definition: 

A number of passengers (1:n) wish to repeatedly take rides with a Roller Coaster. There should be one process for each passengers and one for the Roller Coaster. As for the problem of the sleeping barber, the processes must participate in several synchronization stages. The passengers must wait for the car to be present, and the car must wait for **C** passengers to enter. After the ride, the car must awake all riding passengers. The solution does not have to prevent sneaking. Make a solution in the setting of monitors with signal and continue discipline for signaling. Use the await language to solve this exercise. In addition, try to formulate a reasonable invariant based on your solution, but it is not necessary to prove the invariant using programming logic. 

## Solution:


Needs of the passengers: 
- Passengers must wait until cart is available 
- Passenger must wait may enter the cart if the client

Needs of the cart: 
- Cart must be full before leaving 
- Cart must wake all passengers, and tell them to leave

Cart must wait until:
- Passengers fill all the spots 
- Wait until all passenger leaves

Passengers must wait until:
- Cart arrives
- Cart tells all passengers to leave


```text
monitor RollerCoaster{

    int passengers;               # Represents the amount passengers that are in the cart
    int riding_passengers;        # Represents the passengers when they arrive (when this is set to zero, we know all passengers have left)
    bool cart_ready := false;     # Represent ride being ready to take new passengers
    bool ride_over;               # Represents when a ride is over

    cond cart_arrived;             # Signaled when after all have exited the cart and variables are reset
    cond cart_full;                # Signaled when passengers = C
    cond ride_over;                # Signaled when ride is over and passengers can leave
    cond passengers_left;          # Signal when all riding passengers have left
    

    procedure ride(){
        # Wait until rollercoaster is available, or capacity is full 
        # This will allow sneaking in the queue, but this is okay per assignment description 
        while(!cart_ready || passengers = C){
            wait(cart_arrived);
        }

        # Passenger enters the cart 
        passengers := passengers + 1;

        # If the cart is now full, signal the cart to start the ride
        # A single passenger notifies that the ride is full 
        if (passengers = C){
            signal(cart_full);
        }

        # All passengers ride!

        # Wait until ride is over
        while(!ride_is_over){
            wait(ride_over);
        }

        # After the ride, passenger leaves
        riding_passengers := riding_passengers - 1;

        # Last passenger will signal that there is no more passengers waiting

        if (riding_passengers = 0){
            signal(passengers_left);
        }
    }

    procedure cart_load_passengers(){
        # Wait until all riding passengers on the cart have left
        while(riding_passengers > 0) {wait(passengers_left); }

        # Roller coaster is now ready to accept new passengers
        # Reassign variables to setup a new ride 
        cart_ready := true;
        ride_over := false;
        passengers := 0;
        riding_passengers := C;

        # Notify all passengers that the cart can now except new passengers
        signal_all(cart_arrived);
    }

    procedure cart_unload_passengers(){
        # Signal to all passengers that ride is over
        ride_is_over := true; 
        signal_all(ride_over);
    }
}
```