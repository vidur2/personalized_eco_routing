# Personalized Eco-Routing
## Proof of Concept
### Goals
<ul>
  <li>Calculate hypothetical fuel consumed along a route based on MPG and speed limits</li>
  <li>
    Use a regression model to use user data to predict actual user speed from speed limit
    <ul>
      <li>Regression Model is personalized per individual</li>
    </ul>
  </li>
  <li>
    Calculate Decrease in MPG over a given route based on projected speed
    <ul><li>Select optimal route based on MPG calculation</li></ul>
  </li>
  <li>
    Google Maps-like app on the frontend to select an optimal fuel route and navigate it
  </li>
</ul>

### Current State
This project currently has the backend/api working, but to enable it I need to pay for Google Cloud credit. Once I have tested the API thouroughly, I will being work on a frontend iOS app.

## Full Implementation
### Goals
<ul>
  <li>Get actual fuel data from CarPlay app</li>
  <li>
    Use fuel data to calculate actual fuel consumed based on multiple variables such as:
    <ul>
      <li>
        <strong>Real Time Collected Variables</strong>
        <ul>
          <li>Speed over Speed Limits</li>
          <li>Acceleration</li>
          <li>Gas Consumed</li>
        </ul>
      </li>
      <li>
        <strong>Async Data</strong>
        <ul>
          <li>Traffic Conditions</li>
          <li>Elevation differences</li>
        </ul>
      </li>
    </ul>
  </li>
</ul>
 
 ### Current State 
 Currently not started, making proof of concept first.
