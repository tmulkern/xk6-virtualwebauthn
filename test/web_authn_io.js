
import virtualwebauthn from "k6/x/virtualwebauthn"
import encoding from 'k6/encoding';
import http from 'k6/http';
import { check } from 'k6';


const credentialId="2RLh7QqTDuMdsdv6a62ulw=="
const base64EncodedEC256PkCert="MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgI464z14U9UvD\u002B1A/btftFG\u002BvJPZcGJBvsKNvkJVoJYihRANCAAQnafLhNmzsHdBTaNO/lHNfyPERUjuiUAp1Z87WcJNzM28wdS5fSk/Vw81cTkSXkmlJLgSefhvaSvhwe43iv3cF"

const relyingParty = {
    ID : "webauthn.io",
    Name: "webauthn.io",
    Origin: "https://webauthn.io"
}

const user_id="nobody@test.com"

export default function(){

    http.get("https://webauthn.io/")

    var credential = {
        ID: encoding.b64decode(credentialId),
        KeyType: 0,
        PKCS8Key : encoding.b64decode(base64EncodedEC256PkCert),
        Counter: 0,
    }

    Register(credential);
    
    Login(credential,user_id)

    
}


function Register(credential){
    var attestationOptionsRes=http.post("https://webauthn.io/registration/options",JSON.stringify(
        {
            username:user_id,
            user_verification:"preferred",
            attestation:"none",
            attachment:"all",
            algorithms:[
                "es256",
                "rs256"
            ],
            discoverable_credential:"preferred"
        }),
        {
            headers : {
                "Origin" : "https://webauthn.io",
                "Referer" : "https://webauthn.io"
            }
        });

    check(attestationOptionsRes,{
        "User Registration Options Call Successful" : (r) => r.status === 200
    })

   
    var attestationOptionsJson = attestationOptionsRes.body;

    var attestationOptions=virtualwebauthn.parseAttestationOptions(attestationOptionsJson)
    
    console.log(attestationOptions)
    var authenticator=virtualwebauthn.newAuthenticatorWithOptions({
        UserHandle: toUTF8Array(attestationOptions.user_id).buffer
    })
    

    let responseJson=virtualwebauthn.createAttestationResponse(relyingParty,authenticator,credential,attestationOptions)

    let response = JSON.parse(responseJson)

    response.transports = ["usb"]

    var verificationRequest = JSON.stringify(
        {
            username: user_id,
            authenticatorAttachment: "cross-platform",
            clientExtensionResults : {
                credProps: {
                    rk:false
                }
            },
            response : response
        });

    var verificationResponse=http.post("https://webauthn.io/registration/verification",verificationRequest,
    {
        headers : {
            "Origin" : "https://webauthn.io",
            "Referer" : "https://webauthn.io"
        }
    });

    check(verificationResponse,{
        "User Registration Verification Call Successful" : (r) => r.status === 200
    })
}

function Login(credential,user_id){
    var authenticator=virtualwebauthn.newAuthenticatorWithOptions({
        UserHandle: toUTF8Array(user_id).buffer
    })


    authenticator.addCredential(credential)

    var assertionOptionsResponse=http.post("https://webauthn.io/authentication/options",JSON.stringify(
        {
            username:user_id,
            user_verification:"preferred"
        }),
        {
            headers : {
                "Origin" : "https://webauthn.io",
                "Referer" : "https://webauthn.io"
            }
        });

    check(assertionOptionsResponse,{
        "User Authentication Options Call Successful" : (r) => r.status === 200
    })


    var assertionOptionsJson = assertionOptionsResponse.body


    var assertionOptions = virtualwebauthn.parseAssertionOptions(assertionOptionsJson)

    var assertionResponseJson=virtualwebauthn.createAssertionResponse(relyingParty,authenticator,credential,assertionOptions,true)
    
    var assertionResponse = JSON.parse(assertionResponseJson)

    console.log(assertionResponse)

    assertionResponse.clientExtensionResults={}

    var verificationResponse=http.post("https://webauthn.io/authentication/verification",JSON.stringify(
        {
            username:user_id,
            authenticatorAttachment: "cross-platform",
            response:assertionResponse
        }),
        {
            headers : {
                "Origin" : "https://webauthn.io",
                "Referer" : "https://webauthn.io"
            }
        });

    check(verificationResponse,{
        "User Authentication Verification Call Successful" : (r) => r.status === 200
    })
}

function toUTF8Array(str) {
    var utf8 = [];
    for (var i=0; i < str.length; i++) {
        var charcode = str.charCodeAt(i);
        if (charcode < 0x80) utf8.push(charcode);
        else if (charcode < 0x800) {
            utf8.push(0xc0 | (charcode >> 6), 
                      0x80 | (charcode & 0x3f));
        }
        else if (charcode < 0xd800 || charcode >= 0xe000) {
            utf8.push(0xe0 | (charcode >> 12), 
                      0x80 | ((charcode>>6) & 0x3f), 
                      0x80 | (charcode & 0x3f));
        }
        // surrogate pair
        else {
            i++;
            // UTF-16 encodes 0x10000-0x10FFFF by
            // subtracting 0x10000 and splitting the
            // 20 bits of 0x0-0xFFFFF into two halves
            charcode = 0x10000 + (((charcode & 0x3ff)<<10)
                      | (str.charCodeAt(i) & 0x3ff));
            utf8.push(0xf0 | (charcode >>18), 
                      0x80 | ((charcode>>12) & 0x3f), 
                      0x80 | ((charcode>>6) & 0x3f), 
                      0x80 | (charcode & 0x3f));
        }
    }
    return new Uint8Array(utf8);
}
