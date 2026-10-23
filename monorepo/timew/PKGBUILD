# Maintainer: Jelle van der Waa <jelle@vdwaa.nl>
# Contributor: Joost Bremmer <toost dot b at gmail dot com>

pkgname=timew
pkgver=1.9.1
pkgrel=1
pkgdesc="Timewarrior, A command line time tracking application"
arch=('x86_64')
url="https://timewarrior.net/"
license=('MIT')
makedepends=('cmake' 'gcc-libs' 'asciidoctor')
optdepends=('task: Taskwarrior integration'
            'python: Holiday refresh script')
install=$pkgname.install
source=("https://github.com/GothenburgBitFactory/timewarrior/releases/download/v${pkgver}/timew-${pkgver}.tar.gz")
sha256sums=('7ad34f95c0d61d356df55149f9479f8d9aaec417e5f57f2a1cc76ae2f8a3171b')

build() {
    cd "${pkgname}-${pkgver}"
    cmake -DCMAKE_BUILD_TYPE=None -DCMAKE_INSTALL_PREFIX=/usr .
    make
}


package() {
    cd "${pkgname}-${pkgver}"

    make DESTDIR="${pkgdir}" install

    # Install the refresh holiday scripts in /usr/bin/ for user convenience.
    chmod 755 "${pkgdir}"/usr/share/doc/timew/holidays/refresh
    ln -s /usr/share/doc/timew/holidays/refresh "${pkgdir}"/usr/bin/timew-refresh-holidays

    install -Dm 644 completion/timew-completion.bash "${pkgdir}/usr/share/bash-completion/completions/timew"
    install -Dm 644 completion/timew.zsh "${pkgdir}/usr/share/zsh/site-functions/_timew"
    install -Dm 644 completion/timew.fish "${pkgdir}/usr/share/fish/vendor_completions.d/timew.fish"

    install -Dm 644 LICENSE -t "${pkgdir}/usr/share/licenses/${pkgname}"
}
