import React from 'react'
import firebase from 'firebase'
import StyledFirebaseAuth from 'react-firebaseui/StyledFirebaseAuth'

type SignInWidgetProps = {
  firebaseApp: firebase.app.App
}

const uiConfig = {
  signInFlow: 'popup',
  signInSuccessUrl: '/',
  signInOptions: [
    firebase.auth.GoogleAuthProvider.PROVIDER_ID,
  ],
}

const SignInWidget: React.FC<SignInWidgetProps> = (props: SignInWidgetProps) => {
  return (
    <div>
      <p>Please sign-in:</p>
      <StyledFirebaseAuth uiConfig={uiConfig} firebaseAuth={props.firebaseApp.auth()} />
    </div>
  )
}

export default SignInWidget
