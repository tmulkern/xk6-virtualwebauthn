export function CreateAssertionResponse(relyingParty:RelyingParty, authenticator:Authenticator, portableCredential:PortableCredential, assertionOptions:AssertionOptions):string 

export function CreateAttestationResponse(relyingParty:RelyingParty, authenticator:Authenticator, portableCredential:PortableCredential, attestationOptions:AttestationOptions):string

export function ParseAssertionOptions(jsonString:string):AssertionOptions

export function ParseAttestationOptions(jsonString:string):AttestationOptions

export function NewAuthenticator():Authenticator

export function NewAuthenticatorWithOptions(authenticatorOptions:AuthenticatorOptions):Authenticator


export interface RelyingParty{
    ID:string
    Name:string
    Origin:string
}

export interface PortableCredential{
    ID: ArrayBuffer
    PKCS8Key: ArrayBuffer
    KeyType: number
    Counter: number
}

export interface AttestationOptions {
    Challenge:ArrayBuffer
    ExcludeCredentials:string[]
    RelyingPartyID:string
    RelyingPartyName:string
    UserID:string
    UserName:string
    UserDisplayName:string
}

export interface AuthenticatorOptions {
    UserHandle: ArrayBuffer
    UserNotPresent:boolean
    UserNotVerified:boolean
}


export class Authenticator {
    Options: AuthenticatorOptions
    Aaguid: ArrayBuffer
    Credentials: PortableCredential[]

    AddCredential(portableCredential:PortableCredential): void
    FindAllowedCredential(assertionOptions:AssertionOptions):PortableCredential
}

export interface AssertionOptions {
    Challenge: ArrayBuffer
    AllowCredentials: string[]
    RelyingPartyID: string
}