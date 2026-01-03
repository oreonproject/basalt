{
  lib,
  buildGoModule,
}:

buildGoModule rec {
  pname = "basalt";
  version = "0.1";

  src = ./.;

  vendorHash = "sha256-pprnK2JKmPuR3Q+F8+vMDEdowlb3oX4BOOzW8NGOqgs=";

  ldflags = [ "-s" "-w" ];

  meta = {
    description = "CLI to manage cloud services";
    license = lib.licenses.gpl3Only;
    maintainers = with lib.maintainers; [
      Simon-Weij
      johndavedosn
      taxmalalas0001
    ];
    mainProgram = "basalt";
  };
}
