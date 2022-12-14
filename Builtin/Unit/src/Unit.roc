modeule Unit exposing Length(..), Angle(..), Duration(..)

type Unit
  = Length
  | Angle
  | Duration

{-
  # Length Measurements
-}
type Length
  = Nanometer Float
  | Micrometer Float
  | Millimeter Float
  | Centimeter Float
  | Meter Float
  | Kilometer Float
  | Inch Float
  | Foot Float
  | Yard Float
  | Mile Float

-- entering the type system

nanometers: Float -> Length
nanometers mms = 
  Nanometer mms

micrometers: Float -> Length
micrometers mms = 
  Micrometer mms

millimeters: Float -> Length
millimeters mms = 
  Millimeter mms

centimeters: Float -> Length
centimeters cms =
  Centimeter cms

meters: Float -> Length
meters ms =
  Meter ms
  
kilometers: Float -> Length
kilometers kms =
  Kilometer kms

inches: Float -> Length
inches ins = 
  Inch ins
  
feet: Float -> Length
feet fts = 
  Foot fts

yards: Float -> Length
yards ys = 
  Yard ys

miles: Float -> Length
miles mis = 
  Mile mis
  
  
-- Escaping the type system
inNanometers: Length -> Float
inNanometers len =
  case len of
    Nanometer val -> val
    Micrometer val -> val * 1000
    Millimeter val -> val * (1000^2)
    Centimeter val -> val* 10 * (1000^2)
    Meter val -> val * (1000^3)
    Kilometer val -> val * (1000^4)
    Inch val -> val * 25400000
    Foot val -> val * 12 * 25400000
    Yard val -> val * 3 * 12 * 25400000
    Mile val -> val * 1609000000000


inMicrometers: Length -> Float
inMicrometers len =
  case len of
    Nanometer val -> val / 1000
    Micrometer val -> val
    Millimeter val -> val * 1000
    Centimeter val -> val* 10 * 1000
    Meter val -> val * (1000^2)
    Kilometer val -> val * (1000^3)
    Inch val -> val * 25400
    Foot val -> val * 12 * 25400
    Yard val -> val * 3 * 12 * 25400
    Mile val -> val * 1609000000



inMillimeters: Length -> Float
inMillimeters len =
  case len of
    Nanometer val -> val / (1000^2)
    Micrometer val -> val / 1000
    Millimeter val -> val
    Centimeter val -> val*10
    Meter val -> val * 1000
    Kilometer val -> val * (1000^2)
    Inch val -> val * 25.4
    Foot val -> val * 304.8
    Yard val -> val * 914.4
    Mile val -> val * 1609344 


inCentimeters: Length -> Float
inCentimeters len =
  case len of
    Nanometer val -> val / 10000000
    Micrometer val -> val / 10000
    Millimeter val -> val / 10
    Centimeter val -> val
    Meter val -> val * 100
    Kilometer val -> val * (100000)
    Inch val -> val * 2.54
    Foot val -> val * 30.48
    Yard val -> val * 91.44
    Mile val -> val * 160934


inMeters: Length -> Float
inMeters len =
  case len of
    Nanometer val -> val / (1000^3)
    Micrometer val -> val / (1000^2)
    Millimeter val -> val / 1000
    Centimeter val -> val / 100
    Meter val -> val
    Kilometer val -> val * (1000)
    Inch val -> val * 39.3701
    Foot val -> val * 3.28084
    Yard val -> val * 1.0936133333333
    Mile val -> val * 0.00062137121212119323429

inKilometers: Length -> Float
inKilometers len =
  case len of
    Nanometer val -> val / (1000^4)
    Micrometer val -> val / (1000^3)
    Millimeter val -> val / (1000^2)
    Centimeter val -> val / (100 * 1000)
    Meter val -> val / 1000
    Kilometer val -> val
    Inch val -> val * 39370.1
    Foot val -> val * 3280.84
    Yard val -> val * 1093.61
    Mile val -> val / 0.621371

inInches: Length -> Float
inInches len =
  case len of
    Nanometer val -> val / (2.54e+7)
    Micrometer val -> val / (25400)
    Millimeter val -> val / (25.4)
    Centimeter val -> val / (2.54)
    Meter val -> val * 0.0254
    Kilometer val -> val * 2.54e-5
    Inch val -> val
    Foot val -> val * 12
    Yard val -> val * 36
    Mile val -> val / 63360

inFeet: Length -> Float
inFeet len =
  case len of
    Nanometer val -> val / 3.048e+8
    Micrometer val -> val / 304800
    Millimeter val -> val / 304.8
    Centimeter val -> val / 30.48
    Meter val -> val / 0.3048
    Kilometer val -> val / 0.0003048
    Inch val -> val / 12
    Foot val -> val
    Yard val -> val * 3
    Mile val -> val * 5280

-- I got bored
inYards: Length -> Float
inYards len =
  (inFeet len) / 3

inMiles: Length -> Float
inMiles len =
  (inFeet len) / 5280



-- converting in the type system  

toNanometers: Length -> Length
toNanometers len =
  nanometers (inNanometers len)
  
toMicrometers: Length -> Length
toMicrometers len =
  micrometers (inMicrometers len)

toMillimeters: Length -> Length
toMillimeters len = 
  millimeters (inMillimeters len)

toCentimeters: Length -> Length
toCentimeters len = 
  centimeters (inCentimeters len)

toMeters: Length -> Length
toMeters len = 
  meters (inMeters len)

toKilometers: Length -> Length
toKilometers len = 
  kilometers (inKilometers len)

toInch: Length -> Length
toInch len = 
  inches (inInches len)

toFeet: Length -> Length
toFeet len = 
  feet (inFeet len)
  
toYards: Length -> Length
toYards len = 
  yards (inYards len)

toMiles: Length -> Length
toMiles len = 
  miles (inMiles len)


lengthToString: Length -> String
lengthToString l = 
  case l of
    Nanometer val -> floatWithSuffix val "nm"
    Micrometer val -> floatWithSuffix val "μm"
    Millimeter val -> floatWithSuffix val "mm"
    Centimeter val -> floatWithSuffix val "cm"
    Meter val -> floatWithSuffix val "m"
    Kilometer val -> floatWithSuffix val "km"
    Inch val -> floatWithSuffix val "in"
    Foot val -> floatWithSuffix val "ft"
    Yard val -> floatWithSuffix val "yd"
    Mile val -> floatWithSuffix val "mi"

floatWithSuffix float suffix =
  (String.fromFloat float) ++ suffix


{-
  # Angular Measurements
-}
type Angle 
  = Degree Float
  | Radian Float
  | Revolution Float

-- enterring the type system
degrees: Float -> Angle
degrees dgs = 
  Degree dgs

radians: Float -> Angle
radians rads =
  Radian rads

revolutions: Float -> Angle
revolutions revs =
  Revolution revs
  
-- escaping the type system
inDegrees: Angle -> Float
inDegrees ang =
  case ang of 
    Degree a -> a
    Radian a -> a * 180 / pi
    Revolution a -> a * 360

inRadians: Angle -> Float
inRadians ang =
  case ang of 
    Degree a -> a / 180 * pi
    Radian a -> a
    Revolution a -> a * 2 * pi

inRevolutions: Angle -> Float
inRevolutions ang =
  case ang of
    Degree a -> a /360
    Radian a -> a / (2 * pi)
    Revolution a -> a

-- conversion in the type system
toDegrees: Angle -> Angle
toDegrees ang = 
  degrees (inDegrees ang)

toRadians: Angle -> Angle
toRadians ang = 
  radians (inRadians ang)

toRevolutions ang = 
  revolutions (inRevolutions ang)

angleToString: Angle -> String
angleToString a = 
  case a of
    Degree v -> floatWithSuffix v "degs" 
    Radian v -> floatWithSuffix v "rads"
    Revolution v -> floatWithSuffix v "revs"


{-
  # Duration Measurements
-}
type Duration 
  = Microsecond Float
  | Millisecond Float
  | Second Float
  | Minute Float
  | Hour Float
  | Day Float
  | Week Float


microseconds: Float -> Duration
microseconds ms =
  Microsecond ms

milliseconds: Float -> Duration
milliseconds ms = 
  Millisecond ms

seconds: Float -> Duration
seconds s = 
  Second s

minutes: Float -> Duration
minutes m = 
  Minute m

hours: Float -> Duration
hours h =
  Hour h

days: Float -> Duration
days d = 
  Day d

weeks: Float -> Duration
weeks w =
  Week w

inMicroseconds: Duration -> Float
inMicroseconds dur = 
  case dur of
    Microsecond v -> v
    Millisecond v -> v * 1000
    Second v -> v * (1000^2)
    Minute v -> v * (60 * 1000^2)
    Hour v -> v * (60^2 * 1000^2)
    Day v -> v * (24 * 60^2 * 1000^2)
    Week v -> v * (7 * 24 * 60^2 * 1000^2)

inMilliseconds: Duration -> Float
inMilliseconds dur = 
  case dur of
    Microsecond v -> v / 1000
    Millisecond v -> v 
    Second v -> v * 1000
    Minute v -> v * (60 * 1000)
    Hour v -> v * (60^2 * 1000)
    Day v -> v * (24 * 60^2 * 1000)
    Week v -> v * (7 * 24 * 60^2 * 1000)

inSeconds: Duration -> Float
inSeconds dur = 
  case dur of
    Microsecond v -> v / (1000^2)
    Millisecond v -> v / 1000
    Second v -> v 
    Minute v -> v * 60
    Hour v -> v * (60^2)
    Day v -> v * (24 * 60^2)
    Week v -> v * (7 * 24 * 60^2)

inMinutes: Duration -> Float
inMinutes dur = 
  case dur of
    Microsecond v -> v / (60 * 1000^2)
    Millisecond v -> v / (60 * 1000)
    Second v -> v / 60
    Minute v -> v
    Hour v -> v * 60
    Day v -> v * (24 * 60)
    Week v -> v * (7 * 24 * 60)

inHours: Duration -> Float
inHours dur = 
  case dur of
    Microsecond v -> v / (60^2 * 1000^2)
    Millisecond v -> v / (60^2 * 1000)
    Second v -> v / (60^2)
    Minute v -> v / 60
    Hour v -> v
    Day v -> v * 24
    Week v -> v * (7 * 24)

inDays: Duration -> Float
inDays dur = 
  case dur of
    Microsecond v -> v / (24 * 60^2 * 1000^2)
    Millisecond v -> v / (24 * 60^2 * 1000)
    Second v -> v / (24 * 60^2)
    Minute v -> v / (24 * 60)
    Hour v -> v / 24
    Day v -> v
    Week v -> v * 7

inWeeks: Duration -> Float
inWeeks dur = 
  case dur of
    Microsecond v -> v / (7 * 24 * 60^2 * 1000^2)
    Millisecond v -> v / (7 * 24 * 60^2 * 1000)
    Second v -> v / (7 * 24 * 60^2)
    Minute v -> v / (7 * 24 * 60)
    Hour v -> v / (7 * 24)
    Day v -> v / 7
    Week v -> v

toMicroseconds: Duration -> Duration
toMicroseconds dur = 
  microseconds (inMicroseconds dur)

toMilliseconds: Duration -> Duration
toMilliseconds dur = 
  milliseconds (inMilliseconds dur)

toSeconds: Duration -> Duration
toSeconds dur = 
  seconds (inSeconds dur)

toMinutes: Duration -> Duration
toMinutes dur = 
  minutes (inMinutes dur)

toHours: Duration -> Duration
toHours dur = 
  hours (inHours dur)

toDays: Duration -> Duration
toDays dur = 
  days (inDays dur)

toWeeks: Duration -> Duration
toWeeks dur = 
  weeks (inWeeks dur)

durationToString: Duration -> String
durationToString d = 
  case d of
    Microsecond v -> floatWithSuffix v "μsec" 
    Millisecond v -> floatWithSuffix v "msec"
    Second v -> floatWithSuffix v "sec"
    Minute v -> floatWithSuffix v "min"
    Hour v -> floatWithSuffix v "min"
    Day v -> floatWithSuffix v "days"
    Week v -> floatWithSuffix v "wks"


toString: Unit -> String
toString u ->
  case u of 
    Length -> lengthToString u
    Angle -> angleToString u
    Duration -> durationToString u