import Motor
import DigitalIn

program = 
    { components = 
        { LeftMotor = Motor.new Vex.Port1
        , RightMotor = Motor.new Vex.Port2
        , BumperSwitch = DigitalIn.Input Vex.TriwireA
        }
    , inputs = 
        { DigitalIn.OnChange BumperSwitch
        }
    , outputs = [LeftMotor.SetPercentage, RightMotor.SetPercentage]
    , update = update
    , init = init
}

type Robot
    = Start
    | Stopped

init () -> (Robot, Command)
init = 
    Start, Command.batch []


update: Robot -> Message -> (Robot, Command)
update robot msg =
    case msg of
        ComponentMessage compMsg -> handleComponentInput robot msg


handleComponentInput: Robot -> Message -> (Robot, Command)
handleComponentInput robot compMsg
    case compMsg of
        BumperSwitch.Change value -> 
            if robot == Stopped then
                (Start, Command.batch [LeftMotor.SetPercentage 1, RightMotor.SetPercentage 1])
            else 
                (Stopped, Command.batch [LeftMotor.SetPercentage 0, RightMotor.SetPercentage 0])