# Q1: The Roller Coaster problem

## Problem definition: 

A number of passengers (1:n) wish to repeatedly take rides with a Roller Coaster. There should be one process for each passengers and one for the Roller Coaster. As for the problem of the sleeping barber, the processes must participate in several synchronization stages. The passengers must wait for the car to be present, and the car must wait for C passengers to enter. After the ride, the car must awake all riding passengers. The solution does not have to prevent sneaking. Make a solution in the setting of monitors with signal and continue discipline for signaling. Use the await language to solve this exercise. In addition, try to formulate a reasonable invariant based on your solution, but it is not necessary to prove the invariant using programming logic. 

## Solution:


```text

```
monitor RollerCoaster:
    int waitingPassengers = 0         
    int passengersOnBoard = 0    
    int rollerCoaster = 0     
    int capacity = 0                               

    condition carAvailable  
    condition carLeft            
    condition allPassengersBoarded      
    condition rideFinished              

    procedure carArrives():
        while (rollerCoaster = 0): wait(carAvailable)
        while (passengersOnBoard = capacity): wait(allPassengersBoarded)
        rollerCoaster = rollerCoaster - 1
        signal_all(carLeft) 

    procedure passengerArrives():
        waitingPassengers += 1

        wait(carAvailable)

        passengersOnBoard += 1
        waitingPassengers -= 1

        if passengersOnBoard == capacity:
            signal(allPassengersBoarded)

        wait(rideFinished)