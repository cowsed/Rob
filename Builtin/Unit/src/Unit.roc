modeule Unit exposing Length(..), Angle(..), Time(..)
-- length measurements
type Length 
	= Millimeter Float
	| Centimeter Float
	| Meter Float
	| Inch Float
	| Foot Float

toMillimeter: Length -> Millimeter
toMillimeter length = 
	case dist of 
		Millimeter v -> Millimeter v
		Centimeter -> Millimeter (dist * 10)
		Meter -> Millimeter (dist * 1000)
		Inch -> Millimeter (dist * 25.4)
		Foot -> Millimeter (dist * 12 * 25.4)

toCentimeter: Length -> Centimeter
toCentimeter dist = 
	(toMillimeter dist) * 10
	
toMeter: Length -> Meter
toMeter dist = 
	(toMillimeter dist) * 1000
	
toInch Length -> Inch
toInch dist = 
	case dist of
		Millimeter -> dist / 25.4
		Centimeter -> dist / 2.54
		Meter -> dist / 0.0254
		Inch -> dist 
		Foot -> dist * 12

toFoot dist = 
	(toInch dist) * 12




-- angular measurements

type alias Degree = Float
type alias Radian = Float
type alias Revolution = Float
type Angle 
	= Degree
	| Radian
	| Revolution

toDegree: Angle -> Degree
toDegree ang = 
	case ang of 
		Degree -> Degree
		Radian -> (ang / PI) * 180
		Revolution -> ang * 360

toRadian: Angle -> Radian
toRadian ang = 
	case ang of
		Degree -> (ang / 180) * PI
		Radian -> ang
		Revolution -> ang * 2 * PI

toRevolution: Angle -> Revolution
toRevolution ang = 
	case ang of 
		Degree -> ang / 360
		Radian -> ang / (2 * PI)
		Revolution -> Revolution



-- temporal measurements
type alias Millisecond = Float
type alias Second = Float
type alias Minute = Float
type alias Hour = Float

type Time
	= Millisecond
	| Second
	| Minute
	| Hour

toMillisecond: Time -> Millisecond
toMillisecond time = 
	case Millisecond of
		Millisecond -> time
		Second -> time * 1000
		Minute -> time * 1000 * 60
		Hour -> time * 1000 * 60 * 60

toSecond: Time -> Second
toSecond time = 
	(toMillisecond time) * 1000

toMinute: Time -> Minute
toMinute time = 
	(toSecond time) * 60

toHour: Time -> Hour
toHour time = 
	(toMinute time) * 60
