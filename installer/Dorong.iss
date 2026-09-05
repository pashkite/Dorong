#ifndef MyAppVersion
  #define MyAppVersion "0.6.2"
#endif
#ifndef MyAppExeName
  #define MyAppExeName "Dorong-v0.6.2.exe"
#endif

#define MyAppName "Dorong"
#define MyAppPublisher "pashkite"
#define MyAppURL "https://github.com/pashkite/Dorong"

[Setup]
AppId={{6B3A0C31-8727-4D10-9DA7-6B3B0F132A61}
AppName={#MyAppName}
AppVersion={#MyAppVersion}
AppPublisher={#MyAppPublisher}
AppPublisherURL={#MyAppURL}
AppSupportURL={#MyAppURL}
AppUpdatesURL={#MyAppURL}
DefaultDirName={localappdata}\Programs\Dorong
DefaultGroupName=Dorong
DisableProgramGroupPage=yes
PrivilegesRequired=lowest
OutputDir=..
OutputBaseFilename=Dorong-Setup-v{#MyAppVersion}
Compression=lzma2
SolidCompression=yes
WizardStyle=modern
UninstallDisplayName=Dorong
UninstallDisplayIcon={app}\Dorong.exe
CloseApplications=yes
RestartApplications=no
ArchitecturesAllowed=x64compatible
ArchitecturesInstallIn64BitMode=x64compatible

[Languages]
Name: "english"; MessagesFile: "compiler:Default.isl"
Name: "korean"; MessagesFile: "compiler:Languages\Korean.isl"

[Tasks]
Name: "desktopicon"; Description: "{cm:CreateDesktopIcon}"; GroupDescription: "{cm:AdditionalIcons}"; Flags: unchecked

[Files]
Source: "..\{#MyAppExeName}"; DestDir: "{app}"; DestName: "Dorong.exe"; Flags: ignoreversion

[Icons]
Name: "{group}\Dorong"; Filename: "{app}\Dorong.exe"; WorkingDir: "{app}"
Name: "{autodesktop}\Dorong"; Filename: "{app}\Dorong.exe"; WorkingDir: "{app}"; Tasks: desktopicon

[Run]
Filename: "{app}\Dorong.exe"; Description: "{cm:LaunchProgram,Dorong}"; Flags: nowait postinstall skipifsilent

[UninstallRun]
Filename: "{cmd}"; Parameters: "/C taskkill /IM Dorong.exe /F >NUL 2>&1"; Flags: runhidden; RunOnceId: "StopDorong"
Filename: "{cmd}"; Parameters: "/C reg delete HKCU\Software\Microsoft\Windows\CurrentVersion\Run /v Dorong /f >NUL 2>&1"; Flags: runhidden; RunOnceId: "RemoveDorongStartup"
