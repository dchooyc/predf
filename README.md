# predf

**There are 30 horses. What is the minimum number of races needed so you can identify the fastest 2 horses? You can race up to 5 horses at a time, but you do not have a watch. Please write your answer in detailed steps.**

**For the horses problem:**
In every race we take the fastest two horses
![horses illustration](horses.png)
`go run horse.go`

**There are 127 bottles of water, one of which is poisoned. It takes 3 hours for the poison to take effect. You are given 8 mice. How long will it take to test which bottle of water is poisoned? Please write your answer in detailed steps.**

for the poison bottle problem:
It actually depends. Depends how precious a mouse's life is to you.

**If very precious:**
![slow mouse illustration](slow_mouse.png)

We basically let each mouse test a sip from one different bottle each

We repeat this until a mouse shows a reaction to the poison then that bottle is the poisonous one

This will not take more than ((127/8) + 1) * 3 hrs = 16(Because it was a floor division) * 3 hrs = 48 hrs

This could take shortest just 3 hrs because the first round of tests may show the poison

At most one mouse will be hurt.

**If not precious:**
![fast mouse illustration](fast_mouse.png)

We divide the bottles into almost equal groups (16 * 7) + 15

Each mouse takes a sip from each bottle in a group that is assigned to it

We repeat the step minus the poisoned mouse until we find the poisoned bottle

We use the bottles of the poisoned mouse's group to repeat the step

For each round one mouse will be poisoned

The bottle which the last poisoned mouse sipped from is the poisoned bottle

It will take exactly 9 hrs.

**The company organizes an annual dinner with a total of 801 people. Each person has a
number plate, and each person's number plate is different. Originally, everyone can add
the number plate of another person’s to get a score of 888, but because the total number
of people in the company is odd, there must be someone who did not get a number plate.
Please work out a formula find this single person, and analyze the time taken and tediosity
of the formula.**

Let's say that each of these number plates are an integer in an array of integers

We want to find the index of the number plate without a complementing pair

`go run dinner.go`

The solution's time will be O(n) and the memory usage will also be O(n)
