# Maintainer: Sven-Hendrik Haase <svenstaro@archlinux.org>
# Contributor: t3ddy  <t3ddy1988 "at" gmail {dot} com>
# Contributor: Adrián Chaves Fernández (Gallaecio) <adriyetichaves@gmail.com>
pkgname=0ad
pkgver=0.28.0
pkgrel=3
pkgdesc="Cross-platform, 3D and historically-based real-time strategy game"
arch=('x86_64')
url="http://play0ad.com/"
license=('GPL-2.0-or-later')
depends=("0ad-data=$pkgver" 'binutils' 'boost-libs' 'curl' 'enet' 'libogg' 'libpng' 'libvorbis'
         'libxml2' 'openal' 'sdl2' 'wxwidgets-gtk3' 'zlib' 'libgl' 'glu' 'fmt'
         'gloox' 'miniupnpc' 'libminiupnpc.so' 'icu' 'nspr' 'libsodium' 'which')
makedepends=('boost' 'cmake' 'mesa' 'zip' 'libsm' 'rust' 'git' 'python' 'llvm' 'cbindgen'
             'premake')
options=('!lto') # breaks spidermonkey linking (https://bugs.gentoo.org/746947)
source=("https://releases.wildfiregames.com/$pkgname-$pkgver-unix-build.tar.xz")
validpgpkeys=('A035C8C19219BA821ECEA86B64E628F8D684696D')  # Pablo Galindo Salgado <pablogsal@gmail.com>
sha512sums=('6b5d7893d27d5fa71ce4cad1b424a7d4b36e8bc438c590cbd86d9872148d5051cd3215f8f8f4b817e48d4fe386d8902ecc0dcd612ca779306ef35dfa2898a878')

build() {
  cd "$pkgname-$pkgver"

  # When buidling the included spidermonkey, we need to
  # a compile error: https://gitea.wildfiregames.com/0ad/0ad/issues/8571
  export CXXFLAGS="${CXXFLAGS/-fexceptions/}"

  cd libraries
  ./build-source-libs.sh \
    --with-system-premake

  cd ../build/workspaces
  ./update-workspaces.sh -j$(nproc) \
    --with-system-premake5 \
    --with-lto \
    --without-pch \
    --bindir=/usr/bin \
    --libdir=/usr/lib/0ad \
    --datadir=/usr/share/0ad/data

  cd gcc
  make -j$(nproc)
}

package() {
  cd "$pkgname-$pkgver"
  install -d "${pkgdir}"/usr/{bin,lib/0ad,share/"${pkgname}"/data}
  install -Dm755 binaries/system/pyrogenesis "${pkgdir}/usr/bin"
  install -Dm755 binaries/system/*.so "${pkgdir}/usr/lib/0ad"

  cp -r binaries/data/l10n/ "${pkgdir}/usr/share/${pkgname}/data/"

  install -Dm755 build/resources/${pkgname}.sh "${pkgdir}/usr/bin/${pkgname}"
  install -Dm644 build/resources/${pkgname}.desktop "${pkgdir}/usr/share/applications/${pkgname}.desktop"
  install -Dm644 build/resources/${pkgname}.appdata.xml "${pkgdir}/usr/share/metainfo/${pkgname}.appdata.xml"
  install -Dm644 build/resources/${pkgname}.png "${pkgdir}/usr/share/pixmaps/${pkgname}.png"
}
