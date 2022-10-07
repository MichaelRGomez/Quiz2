module Qelm exposing (main)

import Html exposing (..)
import Html.Attributes exposing (class, src)
import Html.Events exposing (onClick)
import Browser
import Browser.Dom exposing (Viewport)


type Msg = 
    Start | Unstart

initialModel : {start : Bool}
initialModel = {
    start = False
    }

viewIcon : {start : Bool} -> Html Msg
viewIcon model =
    let
        buttonType =
            if model.start then "star" else "star_border"
        msg =
            if model.start then Unstart else Start
    in

    div [class "detialed-icon"] [
        span [class "material-icons md-50", onClick msg ][
            text buttonType
        ]
    ]

update : Msg -> {start : Bool} -> {start : Bool}
update msg model =
    case msg of
        Start -> {model | start = True}
        Unstart -> {model | start = False}

view : {start : Bool} -> Html Msg
view model =
    div [] [
        div [class "content-flow"] [
            viewIcon model
        ]
    ]

main : Program () {start : Bool} Msg
main =
    Browser.sandbox{
        init = initialModel
        ,
        view = view
        ,
        update = update
    }