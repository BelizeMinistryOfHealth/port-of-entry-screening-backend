# Port of Entry COVID10 Screening App
The Backend of the COVID19 Port Of Entry Screening App

## Data Model
The data model can be subdivided into 4 general areas:
1. Arrival Info
  - Date of Arrival
  - Mode of Arrival (Air/Land-Bus/Land-Private Vehicle/Water - Boat)
  - Vessel Number
  - City_Airport (of origin)
  - Port of Entry in BZ
2. Passenger Info
  - First Name
  - Middle Name Initial
  - Last Name
  - Date of Birth
  - Nationality
  - Home Address
  - Home District/Province
  - Origin of Travel
  - Last port of Embarkation
  - Date of Embarkation
3. Address In Belize
  - District
  - City_Town_Village
  - Address
  - Travelling Companions: [{ CompanionId, Relationship?}]
4. Screening
  - Flu Like Symptoms: [{ Symptom, Bool}]
  - Other Symptoms
  - Diagnosed with COVID19 
  - Contact With COVID19
  - Contact with Health Facility
  - Comments
  - Location Person was screened 

### Gaps
[] How do we store meta information? For example:
  - Who entered the information.
  - Where it was entered. This particular information can make syncing more efficient.
    As the information grows, it would be impractical to download all the data into the tablets.
    We might want to be selective about doing this, and only download the information related to
    the location where the tablet is (or will be).

 

