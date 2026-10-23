# Maintainer: Sven-Hendrik Haase <svenstaro@archlinux.org>
# Contributor: javier dot tia at gmail dot com
# Contributor: Paul Davis ("dangersalad") <paul@dangersalad.com>

pkgname=ttf-input-nerd
pkgver=3.5.1
_ttf_input_pkgver=20220502
pkgrel=1
pkgdesc='Patched font Input containing nerd font symbols'
arch=('any')
url='https://input.djr.com/'
license=('custom')
makedepends=('fontforge' 'parallel')
provides=('ttf-font' 'ttf-font-nerd')
_url="${url}build/?fontSelection=whole&a=0&g=0&i=0&l=0&zero=0&asterisk=0&braces=0&preset=default&line-height=1.2&accept=I+do&email="
# the content at $_url is dynamically generated every request, use a snapshot
source=("https://sources.archlinux.org/other/packages/ttf-input/Input_Fonts_${_ttf_input_pkgver}.zip"
        nerd-fonts-patcher-$pkgver.zip::"https://github.com/ryanoasis/nerd-fonts/releases/download/v${pkgver}/FontPatcher.zip")
sha256sums=('bc314ea9026fba5134cd2a7ba93a4ec48607f341d32e4eb3dfae20d141cb91b3'
            '42bcb32145499a35732274c7fc48deb434ad0d2e0e118f98527c1479c6fa251a')

build () {
  mkdir -p ${srcdir}/TTF
  find ${srcdir}/Input_Fonts -name '*.ttf' \
       | parallel fontforge -script ${srcdir}/font-patcher \
       --no-progressbars --careful \
       --complete --outputdir \
       ${srcdir}/TTF {} \;
}

package() {
  mkdir -p ${pkgdir}/usr/share/fonts/TTF
  cp -r ${srcdir}/TTF/* ${pkgdir}/usr/share/fonts/TTF
  chmod 644 ${pkgdir}/usr/share/fonts/TTF/*

  install -Dm644 "${srcdir}/LICENSE.txt" "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE.txt"
}
